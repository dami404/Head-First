package prose

import (
	"fmt"
	"testing"
)

type JoinWithCommasData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	testCases := []JoinWithCommasData{
		{list: []string{}, want: ""},
		{list: []string{"a"}, want: "a"},
		{list: []string{"a", "b"}, want: "a and b"},
		{list: []string{"a", "b", "c"}, want: "a, b, and c"},
	}

	for _, tc := range testCases {
		actualResult := JoinWithCommas(tc.list)
		if tc.want != actualResult {
			printGotWant(actualResult, tc.want)
		}
	}
}

func printGotWant(actual, wanted string) string {
	return fmt.Sprintf("GOT: %v | WANT: %v", actual, wanted)
}
