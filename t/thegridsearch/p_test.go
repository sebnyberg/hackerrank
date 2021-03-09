package thegridsearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_gridSearch(t *testing.T) {
	for _, tc := range []struct {
		G    []string
		P    []string
		want string
	}{
		{
			[]string{
				"1234567890",
				"0987654321",
				"1111111111",
				"1111111111",
				"2222222222",
			}, []string{
				"876543",
				"111111",
				"111111",
			},
			"YES",
		},
		{
			[]string{
				"123456",
				"567890",
				"234567",
				"194729",
			}, []string{
				"2345",
				"6789",
				"3456",
				"9472",
			}, "YES",
		},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.G, tc.P), func(t *testing.T) {
			require.Equal(t, tc.want, gridSearch(tc.G, tc.P))
		})
	}
}

// Complete the gridSearch function below.
func gridSearch(G []string, P []string) string {
	g := make([][]byte, len(G))
	p := make([][]byte, len(P))
	for i := range g {
		g[i] = []byte(G[i])
	}
	for i := range p {
		p[i] = []byte(P[i])
	}
	pn := len(p[0])
	pm := len(p)
	gn := len(g[0])
	gm := len(g)
	for i := 0; i < gm-pm+1; i++ {
		for j := 0; j < gn-pn+1; j++ {
			for k := 0; k < pm; k++ {
				for l := 0; l < pn; l++ {
					if g[i+k][j+l] != p[k][l] {
						goto ContinueSearch
					}
				}
			}
			return "YES"
		ContinueSearch:
		}
	}
	return "NO"
}
