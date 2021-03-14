package sherlockandvalidstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isValid(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd", "YES"},
		{"aaaaabc", "NO"},
		{"abcdefghhgfedecba", "YES"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isValid(tc.s))
		})
	}
}

func isValid(s string) string {
	var chs [26]int
	for _, ch := range s {
		chs[ch-'a']++
	}
	nCharsWithCount := make(map[int]int)
	for _, count := range chs {
		if count > 0 {
			nCharsWithCount[count]++
		}
	}
	if len(nCharsWithCount) == 1 {
		return "YES"
	} else if len(nCharsWithCount) > 2 {
		return "NO"
	}
	var counts [2]int
	var i int
	for k := range nCharsWithCount {
		counts[i] = k
		i++
	}
	if nCharsWithCount[counts[1]] > nCharsWithCount[counts[0]] {
		counts[0], counts[1] = counts[1], counts[0]
	}
	// counts[0] has the max number of chars
	switch {
	case nCharsWithCount[counts[1]] > 1,
		nCharsWithCount[counts[1]] == 1 && abs(counts[1]-counts[0]) != 1 && counts[1] != 1:
		return "NO"
	default:
		return "YES"
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
