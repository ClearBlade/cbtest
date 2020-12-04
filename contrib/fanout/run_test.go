package fanout

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRun_LogAndLogfProxySucceeds(t *testing.T) {

	words := []string{"foo", "bar", "baz"}

	// mockT expectations

	mockT := &mocks.T{}
	mockT.On("Helper").Return()

	mockT.On("Log", "fanout/Log_words/0:", "Received word:", "foo")
	mockT.On("Log", "fanout/Log_words/1:", "Received word:", "bar")
	mockT.On("Log", "fanout/Log_words/2:", "Received word:", "baz")

	mockT.On("Logf", "%s Received word: %s", "fanout/Log_words/0:", "foo")
	mockT.On("Logf", "%s Received word: %s", "fanout/Log_words/1:", "bar")
	mockT.On("Logf", "%s Received word: %s", "fanout/Log_words/2:", "baz")

	// test

	runJob := Run(mockT, "Log words", len(words), func(m cbtest.T, idx int) {
		m.Log("Received word:", words[idx])
		m.Logf("Received word: %s", words[idx])
	})

	runJob.Wait()

	// assert

	mockT.AssertExpectations(t)
}

func TestRun_ChannelCommunicationSucceeds(t *testing.T) {

	words := []string{"foo", "bar", "baz"}

	mockT := &mocks.T{}
	mockT.On("Helper").Return()

	c := make(chan string)
	m := sync.Map{}

	Run(mockT, "Send words", len(words), func(t cbtest.T, idx int) {
		c <- words[idx]
	})

	receiver := Run(mockT, "Receive words", len(words), func(t cbtest.T, idx int) {
		received := <-c
		m.Store(received, struct{}{})
	})

	receiver.Wait()

	for _, word := range words {
		_, ok := m.Load(word)
		assert.True(t, ok, "has %s key", word)
	}
}

func TestRun_ConsumersProducersSucceeds(t *testing.T) {

	consumers := []string{"consumer-0", "consumer-1"}
	producers := []string{"producer-0", "producer-1", "producer-2", "producer-3"}
	perProducer := 10

	// mockT expectations

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Logf", mock.Anything, mock.Anything, mock.Anything).Return()

	// execute test

	c := make(chan int)
	totalProduced := uint32(0)
	totalConsumed := uint32(0)

	consumeJob := Run(mockT, "Consumers", len(consumers), func(t cbtest.T, idx int) {

		t.Logf("%s started", consumers[idx])

		for num := range c {
			t.Logf("Consumed: %d", num)
			atomic.AddUint32(&totalConsumed, 1)
		}
	})

	produceJob := Run(mockT, "Producers", len(producers), func(t cbtest.T, idx int) {

		t.Logf("%s started", producers[idx])

		for idx := 0; idx < perProducer; idx++ {
			c <- idx
			t.Logf("Produced: %d", idx)
			atomic.AddUint32(&totalProduced, 1)
		}
	})

	produceJob.Wait()
	close(c)
	consumeJob.Wait()

	// assertions

	assert.Equal(t, totalProduced, totalConsumed, "Produced equals consumed")
}
