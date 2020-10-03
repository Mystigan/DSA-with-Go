package leetcode

import "testing"

var testcases = []struct {
	Input      string
	SearchWord string
	Output     int
}{
	{
		Input:      "i love eating burger",
		SearchWord: "burg",
		Output:     4,
	},
	{
		Input:      "this problem is an easy problem",
		SearchWord: "pro",
		Output:     2,
	},
	{
		Input:      "i am tired",
		SearchWord: "you",
		Output:     -1,
	},
	{
		Input:      "i use triple pillow",
		SearchWord: "pill",
		Output:     4,
	},
	{
		Input:      "hello from the other side",
		SearchWord: "they",
		Output:     -1,
	},
}

func TestIsPrefixOfWord(t *testing.T) {
	for _, tc := range testcases {
		got := isPrefixOfWord(tc.Input, tc.SearchWord)
		want := tc.Output
		if got != want {
			t.Errorf("got: %v, want: %v, input: %v", got, want, tc.Input)
		}
	}
}
