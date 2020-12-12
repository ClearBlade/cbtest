// Package flow showcases a test with a more complex test structure using the
// flow module. It consists of three sequential stages, with most sub-tests in
// each stage running in parallel.
package flow

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	cb "github.com/clearblade/Go-SDK"
	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/contrib/flow"
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

// Main

func TestFlow(t *testing.T) {

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

	// run workflow against the created system
	runWorkflow(t, s)
}

// Workflow

func runWorkflow(t *testing.T, s *system.EphemeralSystem) {

	// total messages published (shared global between workers)
	totalPublished := int32(0)

	// device context memoizer (separate context per worker)
	devices := flow.NewMemoizer()

	// build workflow
	workflow := flow.NewBuilder().Sequence(func(b *flow.Builder) {

		// Stage 0 (Initialize)

		b.WithName("Devices create and connect").Parallel(func(b *flow.Builder) {
			for idx := 0; idx < *flagDevices; idx++ {
				b.WithContext(devices.Get(idx)).Run(deviceInit(s))
			}
		})

		// Stage 1 (Publish)

		b.WithName("Devices publish").Parallel(func(b *flow.Builder) {
			for idx := 0; idx < *flagDevices; idx++ {
				b.WithContext(devices.Get(idx)).Run(devicePublish(s, &totalPublished))
			}
		})

		// Stage 2 (Check)

		b.WithName("Check results").Run(checkResults(s, &totalPublished))
	})

	// run workflow (test will fail if workflow fails)
	flow.Run(t, workflow)
}

func deviceInit(s *system.EphemeralSystem) flow.Worker {

	return func(t *flow.T, ctx flow.Context) {

		idx := ctx.Identifier()
		name := fmt.Sprintf("Device-%d", idx)

		t.Log("Registering device...")
		auth.RegisterDevice(t, s, name, "active-key")

		t.Log("Logging in device...")
		deviceClient := auth.LoginDevice(t, s, name, "active-key")

		t.Log("Initializing MQTT...")
		mqtt.InitializeMQTT(t, s, deviceClient)

		// stash values that we can use later
		ctx.Stash("deviceClient", deviceClient)
	}
}

func devicePublish(s *system.EphemeralSystem, totalPublished *int32) flow.Worker {

	return func(t *flow.T, ctx flow.Context) {

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
			atomic.AddInt32(totalPublished, 1)

			// sleep until we can publish again
			time.Sleep(*flagPeriod)
		}
		t.Log("Publishing done")
	}
}

func checkResults(s *system.EphemeralSystem, totalPublished *int32) flow.Worker {

	return func(t *flow.T, ctx flow.Context) {

		// ID of the collection we are gonna check
		collID := collection.IDByName(t, s, MessagesCollection)

		// total rows in the collection
		totalRows := collection.Total(t, s, collID)

		// check the number of rows equals number of messages published
		should.ExpectE(t, totalRows, to.BeNumerically("==", *totalPublished), "Collection rows does not match total messages published")

		// show results
		logResults(t, totalRows, int(*totalPublished))
	}
}

// Log helpers

func logFlags(t cbtest.T) {
	t.Helper()
	t.Logf("Publish duration: %s", *flagDuration)
	t.Logf("Service instances running: %d", *flagInstances)
	t.Logf("Total devices publishing: %d", *flagDevices)
	t.Logf("Period between publishes: %s", *flagPeriod)
}

func logResults(t cbtest.T, totalRows, totalPublished int) {
	t.Helper()
	t.Logf("Publish duration: %s", *flagDuration)
	t.Logf("Messages published: %d", totalPublished)
	t.Logf("Messages in collection: %d", totalRows)
	t.Logf("Messages per/sec: %d", totalRows/int(flagDuration.Seconds()))
}
