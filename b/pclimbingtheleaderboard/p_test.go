package pclimbingtheleaderboard

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_climbingLeaderboard(t *testing.T) {
	for _, tc := range []struct {
		ranked []int32
		player []int32
		want   []int32
	}{
		{
			[]int32{100, 90, 90, 80, 75, 60},
			[]int32{50, 65, 77, 90, 102},
			[]int32{6, 5, 4, 2, 1},
		},
		{
			[]int32{100, 100, 50, 40, 40, 20, 10},
			[]int32{5, 25, 50, 120},
			[]int32{6, 4, 2, 1},
		},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.ranked, tc.player), func(t *testing.T) {
			require.Equal(t, tc.want, climbingLeaderboard(tc.ranked, tc.player))
		})
	}
}

type RankNode struct {
	points int32
	next   *RankNode
}

func climbingLeaderboard(ranked []int32, player []int32) (results []int32) {
	dummy := &RankNode{
		points: math.MinInt32,
	}
	cur := dummy
	n := 0
	for i := len(ranked) - 1; i >= 0; i-- {
		r := ranked[i]
		for r > cur.points {
			if cur.next == nil {
				cur.next = &RankNode{
					points: r,
				}
				n++
			}
			cur = cur.next
		}
	}

	cur = dummy
	curRank := n
	for _, p := range player {
		for cur.next != nil && p > cur.next.points {
			curRank--
			cur = cur.next
		}
		switch {
		case cur.next == nil:
			results = append(results, 1)
		case p < cur.next.points:
			curRank++
			cur.next = &RankNode{
				points: p,
				next:   cur.next,
			}
			results = append(results, int32(curRank))
		case p == cur.next.points:
			results = append(results, int32(curRank))
		}
	}
	return results
}
