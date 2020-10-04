package leetcode

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output []string
}{
	{
		Input:  []string{"mass", "as", "hero", "superhero"},
		Output: []string{"as", "hero"},
	},
	{
		Input:  []string{"leetcode", "et", "code"},
		Output: []string{"et", "code"},
	},
	{
		Input:  []string{"blue", "green", "bu"},
		Output: []string{},
	},
}

func TestStringMatching(t *testing.T) {
	for _, tc := range testCases {
		want := tc.Output
		got := stringMatching(tc.Input)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}
