package acmicpcteam

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_acmTeam(t *testing.T) {
	for _, tc := range []struct {
		topic []string
		want  []int32
	}{
		{[]string{"10101", "11110", "00010"}, []int32{5, 1}},
		{[]string{"10101", "11100", "11010", "00101"}, []int32{5, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.topic), func(t *testing.T) {
			require.Equal(t, tc.want, acmTeam(tc.topic))
		})
	}
}

// Complete the acmTeam function below.
func acmTeam(topic []string) []int32 {
	teamTopicsCount := make(map[int32]int32)
	var maxTopics int32
	for i := 0; i < len(topic); i++ {
		for j := i + 1; j < len(topic); j++ {
			teamTopics := countTopics(topic[i], topic[j])
			if teamTopics > maxTopics {
				maxTopics = teamTopics
			}
			teamTopicsCount[teamTopics]++
		}
	}
	return []int32{maxTopics, teamTopicsCount[maxTopics]}
}

func countTopics(first, second string) (res int32) {
	for i := range first {
		if first[i] == '1' || second[i] == '1' {
			res++
		}
	}
	return res
}
