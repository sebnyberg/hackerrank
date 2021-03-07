package organizingcontainers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_organizingContainers(t *testing.T) {
	for _, tc := range []struct {
		container [][]int32
		want      string
	}{
		{[][]int32{{1, 3, 1}, {2, 1, 2}, {3, 3, 3}}, "Impossible"},
		{[][]int32{{0, 2, 1}, {1, 1, 1}, {2, 0, 0}}, "Possible"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.container), func(t *testing.T) {
			require.Equal(t, tc.want, organizingContainers(tc.container))
		})
	}
}

func organizingContainers(container [][]int32) string {
	// We want to organize the balls in such a way that each container
	// contains only one kind of ball
	// Container capacity will always remain the same, which means
	// that whether or not it is possible to organize container is dependent
	// on whether each capacity is assignable to the total number of balls
	// per type
	capacities := make([]int32, len(container))
	types := make([]int32, len(container))
	for i := range container {
		for j := range container[i] {
			capacities[i] += container[i][j]
			types[j] += container[i][j]
		}
	}

	capacityCount := make(map[int32]int32)
	for _, c := range capacities {
		capacityCount[c]++
	}

	for _, t := range types {
		if capacityCount[t] > 0 {
			capacityCount[t]--
		} else {
			return "Impossible"
		}
	}

	return "Possible"
}
