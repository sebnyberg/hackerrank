package separatethenumbers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}

func Test_separateNumbers(t *testing.T) {
	// for _, tc := range []struct {
	// 	s string
	// }{
	// 	{"99100"},
	// 	{"91011"},
	// 	{"1234"},
	// } {
	// 	t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
	// 		separateNumbers2(tc.s)
	// 	})
	// }

	t.Run("with input", func(t *testing.T) {
		inPath := "testdata/input"
		wantPath := "testdata/want"
		f, _ := os.Open(inPath)
		defer f.Close()
		sc := bufio.NewScanner(f)
		sc.Scan()
		n := MustParseInt(sc.Text())
		inputs := make([]string, n)
		for i := 0; sc.Scan(); i++ {
			inputs[i] = sc.Text()
		}
		wantFile, _ := os.Open(wantPath)
		defer wantFile.Close()
		sc = bufio.NewScanner(wantFile)
		want := make([]string, n)
		for i := 0; sc.Scan(); i++ {
			want[i] = sc.Text()
		}
		for i := range inputs {
			got := separateNumbers2(inputs[i])
			require.Equal(t, want[i], got)
		}
	})
}

func separateNumber(s string) {
	fmt.Println(separateNumbers2(s))
}

// Complete the separateNumbers function below.
func separateNumbers2(s string) string {
	shiftAt := 1
	for startWidth := 1; startWidth <= len(s)/2; startWidth++ {
		n, err := strconv.Atoi(s[:startWidth])
		if err != nil {
			log.Fatalln(err)
		}
		curWidth := startWidth
		shiftAt *= 10
		prev := n
		nextShiftAt := shiftAt
		for start := curWidth; ; start += curWidth {
			if prev+1 == nextShiftAt {
				curWidth++
				nextShiftAt *= 10
			}
			if start+curWidth > len(s) {
				break
			}
			if s[start] == '0' {
				break
			}
			cur, err := strconv.Atoi(s[start : start+curWidth])
			if err != nil {
				log.Fatalln(err)
			}
			if cur != prev+1 {
				break
			}
			if start+curWidth >= len(s) {
				return fmt.Sprintf("YES %v", n)
			}
			prev = cur
		}
	}
	return "NO"
}
