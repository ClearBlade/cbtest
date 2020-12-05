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
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/contrib/fanout"
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

	// stage-0

	devicesCreateAndConnect := fanout.Run(t, "Devices create and connect", *flagDevices, stageCreateAndConnect(s))

	fanout.Wait(t, devicesCreateAndConnect, time.Minute)

	// stage-1

	totalPublishes := int32(0)

	devicesPublish := fanout.Continue(t, "Devices publish", devicesCreateAndConnect, stagePublish(s, &totalPublishes))

	fanout.Wait(t, devicesPublish, time.Minute)

	// check results

	checkResults(t, s, int(totalPublishes))
}

// Stages

func stageCreateAndConnect(s *system.EphemeralSystem) fanout.WorkerFunc {
	return func(t cbtest.T, ctx fanout.Context) {

		idx := ctx.Identifier()
		name := fmt.Sprintf("Device-%d", idx)

		t.Log("Registering device")
		auth.RegisterDevice(t, s, name, "active-key")

		t.Log("Logging in")
		deviceClient := auth.LoginDevice(t, s, name, "active-key")

		t.Log("Initializing MQTT")
		mqtt.InitializeMQTT(t, s, deviceClient)

		// stash values that we can use later
		ctx.Stash("deviceClient", deviceClient)
	}
}

func stagePublish(s *system.EphemeralSystem, totalPublishes *int32) fanout.WorkerFunc {
	return func(t cbtest.T, ctx fanout.Context) {

		// unstash client from previous stage
		deviceClient, ok := ctx.Unstash("deviceClient").(*cb.DeviceClient)
		should.Expect(t, ok, to.BeTrue())

		t.Log("Publishing...")
		start := time.Now()
		for time.Since(start) < *flagDuration {

			// generate message to send
			message := GenerateMessage()
			data, err := json.Marshal(message)
			should.NoError(t, err)

			// send message
			err = deviceClient.Publish(MessagesTopic, data, 1)
			should.NoError(t, err)

			// increment global counter using atomic operation
			atomic.AddInt32(totalPublishes, 1)

			// sleep until we can publish again
			time.Sleep(*flagPeriod)
		}
		t.Log("Publishing done")
	}
}

// Test helpers

func logFlags(t *testing.T) {
	t.Logf("Publish duration: %s", *flagDuration)
	t.Logf("Service instances running: %d", *flagInstances)
	t.Logf("Total devices publishing: %d", *flagDevices)
	t.Logf("Period between publishes: %s", *flagPeriod)
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
