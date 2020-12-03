package fanout

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumSeries(n int) int {
	return n * (n + 1) / 2
}

func TestRunSucceeds(t *testing.T) {

	tests := []int{
		0,
		1,
		2,
		4,
		8,
		16,
		32,
		64,
		128,
		256,
		512,
		1024,
	}

	for _, tt := range tests {

		name := fmt.Sprintf("Parallel %d", tt)

		t.Run(name, func(t *testing.T) {

			sumTotal := int32(0)

			Run(t, name, tt, func(t *testing.T, idx int) {
				atomic.AddInt32(&sumTotal, int32(idx+1))
			})

			assert.Equal(t, sumSeries(tt), int(sumTotal), "Sum with formula and atomic sum in workers")
		})
	}
}
