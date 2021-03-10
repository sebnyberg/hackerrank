package caesarcipher

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_caesarCipher(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int32
		want string
	}{
		{"There's-a-starman-waiting-in-the-sky", 3, "Wkhuh'v-d-vwdupdq-zdlwlqj-lq-wkh-vnb"},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.s, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, caesarCipher(tc.s, tc.k))
		})
	}
}

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	res := make([]byte, 0)
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, r := range s {
		if strings.ContainsRune(lower, r) {
			r = ((r - 'a' + k) % 26) + 'a'
		} else if strings.ContainsRune(upper, r) {
			r = ((r - 'A' + k) % 26) + 'A'
		}
		res = append(res, byte(r))
	}
	return string(res)
}
