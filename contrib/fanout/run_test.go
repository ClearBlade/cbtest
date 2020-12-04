package fanout_test

import (
	"fmt"
	"testing"

	"github.com/clearblade/cbtest"
	"github.com/clearblade/cbtest/contrib/fanout"
	"github.com/clearblade/cbtest/mocks"
)

func TestRun_LogAndLogfProxySucceeds(t *testing.T) {

	// mockT expectations

	mockT := &mocks.T{}
	mockT.On("Helper").Return()
	mockT.On("Log", "fanout/Create_words/0:", "Created word:", "word-0")
	mockT.On("Log", "fanout/Create_words/1:", "Created word:", "word-1")
	mockT.On("Log", "fanout/Create_words/2:", "Created word:", "word-2")
	mockT.On("Logf", "%s Created word: %s", "fanout/Create_words/0:", "word-0")
	mockT.On("Logf", "%s Created word: %s", "fanout/Create_words/1:", "word-1")
	mockT.On("Logf", "%s Created word: %s", "fanout/Create_words/2:", "word-2")
	mockT.On("Logf", "Waiting for group \"%s\"...", "Create words")

	// test

	createWords := fanout.Run(mockT, "Create words", 3, func(t cbtest.T, ctx fanout.Context) {
		word := fmt.Sprintf("word-%d", ctx.Identifier())
		t.Log("Created word:", word)
		t.Logf("Created word: %s", word)
	})

	fanout.Wait(mockT, createWords)

	// assert

	mockT.AssertExpectations(t)
}

// func TestRun_ChannelCommunicationSucceeds(t *testing.T) {

// 	words := []string{"foo", "bar", "baz"}

// 	mockT := &mocks.T{}
// 	mockT.On("Helper").Return()

// 	c := make(chan string)
// 	m := sync.Map{}

// 	Run(mockT, "Send words", len(words), func(t cbtest.T, idx int) {
// 		c <- words[idx]
// 	})

// 	receiver := Run(mockT, "Receive words", len(words), func(t cbtest.T, idx int) {
// 		received := <-c
// 		m.Store(received, struct{}{})
// 	})

// 	receiver.Wait()

// 	for _, word := range words {
// 		_, ok := m.Load(word)
// 		assert.True(t, ok, "has %s key", word)
// 	}
// }

// func TestRun_ConsumersProducersSucceeds(t *testing.T) {

// 	consumers := []string{"consumer-0", "consumer-1"}
// 	producers := []string{"producer-0", "producer-1", "producer-2", "producer-3"}
// 	perProducer := 10

// 	// mockT expectations

// 	mockT := &mocks.T{}
// 	mockT.On("Helper").Return()
// 	mockT.On("Logf", mock.Anything, mock.Anything, mock.Anything).Return()

// 	// execute test

// 	c := make(chan int)
// 	totalProduced := uint32(0)
// 	totalConsumed := uint32(0)

// 	consumeJob := Run(mockT, "Consumers", len(consumers), func(t cbtest.T, idx int) {

// 		t.Logf("%s started", consumers[idx])

// 		for num := range c {
// 			t.Logf("Consumed: %d", num)
// 			atomic.AddUint32(&totalConsumed, 1)
// 		}
// 	})

// 	produceJob := Run(mockT, "Producers", len(producers), func(t cbtest.T, idx int) {

// 		t.Logf("%s started", producers[idx])

// 		for idx := 0; idx < perProducer; idx++ {
// 			c <- idx
// 			t.Logf("Produced: %d", idx)
// 			atomic.AddUint32(&totalProduced, 1)
// 		}
// 	})

// 	produceJob.Wait()
// 	close(c)
// 	consumeJob.Wait()

// 	// assertions

// 	assert.Equal(t, totalProduced, totalConsumed, "Produced equals consumed")
// }
