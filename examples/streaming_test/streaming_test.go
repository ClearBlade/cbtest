// Package streatming showcases a test that interacts and measures simple
// metrics for a streaming service.
// See: https://blog.golang.org/subtests
package streaming

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest/modules/auth"
	"github.com/clearblade/cbtest/modules/collection"
	"github.com/clearblade/cbtest/modules/mqtt"
	"github.com/clearblade/cbtest/modules/should"
	"github.com/clearblade/cbtest/modules/should/to"
	"github.com/clearblade/cbtest/modules/system"
)

// Globals

const (
	StreamLoggerService = "streamLogger"
	MessagesTopic       = "messages"
	MessagesCollection  = "messages"
)

// Flags

var (
	flagDuration  = flag.Duration("duration", time.Second*5, "The duration of the publish")
	flagInstances = flag.Int("instances", 1, "Stream service instances")
	flagDevices   = flag.Int("devices", 10, "Total devices that are gonna be publishing messages")
	flagPeriod    = flag.Duration("period", time.Millisecond*100, "Period between publishes")
)

// Test

func TestStreaming(t *testing.T) {

	logFlags(t)

	// import into new system
	s := system.UseOrImport(t, "./extra")

	// close the system after the test
	defer system.Finish(t, s)

	// obtain developer client from the ephemeral system
	devClient := auth.LoginAsDev(t, s)

	// start the service
	err := devClient.SetLongRunningServiceParams(s.SystemKey(), StreamLoggerService, false, true, *flagInstances)
	should.NoError(t, err)

	// ID of the collection
	collID := collection.IDByName(t, s, MessagesCollection)

	// clears messages collection
	err = devClient.DeleteData(collID, cb.NewQuery())
	should.NoError(t, err)

	// connect each device and publish
	totalPublishes := devicesConnectAndPublish(t, s)

	// check results
	checkResults(t, s, totalPublishes)
}

// Test helpers

func logFlags(t *testing.T) {
	t.Logf("Publish duration: %s", *flagDuration)
	t.Logf("Service instances running: %d", *flagInstances)
	t.Logf("Total devices publishing: %d", *flagDevices)
	t.Logf("Period between publishes: %s", *flagPeriod)
}

func devicesConnectAndPublish(t *testing.T, s *system.EphemeralSystem) int {

	var totalPublishes int32

	// NOTE: will block until all workers are done.
	// see: https://blog.golang.org/subtests#TOC_7.
	t.Run("Devices connect and publish", func(t *testing.T) {
		for idx := 0; idx < *flagDevices; idx++ {
			name := fmt.Sprintf("Device-%d", idx)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				workerPublishes := deviceWorker(t, s, name)
				atomic.AddInt32(&totalPublishes, int32(workerPublishes))
			})
		}
	})

	return int(totalPublishes)
}

func deviceWorker(t *testing.T, s *system.EphemeralSystem, name string) int {

	messagesPublished := 0

	// register device
	auth.RegisterDevice(t, s, name, "active-key")

	// login as the device we just created
	deviceClient := auth.LoginDevice(t, s, name, "active-key")
	t.Log("Logged in")

	// initialize mqtt
	mqtt.InitializeMQTT(t, s, deviceClient)

	// publish
	start := time.Now()
	for time.Since(start) < *flagDuration {

		// generate message to send
		message := GenerateMessage()
		data, err := json.Marshal(message)
		should.NoError(t, err)

		// send message
		err = deviceClient.Publish(MessagesTopic, data, 1)
		should.NoError(t, err)
		messagesPublished++

		// sleep until we can publish again
		time.Sleep(*flagPeriod)
	}

	t.Log("Done publishing")
	return messagesPublished
}

func checkResults(t *testing.T, s *system.EphemeralSystem, messagesPublished int) {
	t.Run("Check results", func(t *testing.T) {

		// ID of the collection we are gonna check
		collID := collection.IDByName(t, s, MessagesCollection)

		// total rows in the collection
		totalRows := collection.Total(t, s, collID)

		// check the number of rows equals number of messages published
		should.ExpectE(t, totalRows, to.Equal(messagesPublished), "Collection rows does not match total messages published")

		// logs results
		t.Logf("Publish duration: %s", *flagDuration)
		t.Logf("Messages published: %d", messagesPublished)
		t.Logf("Messages in collection: %d", totalRows)
		t.Logf("Messages per/sec: %d", totalRows/int(flagDuration.Seconds()))
	})
}
