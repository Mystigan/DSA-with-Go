package leetcode

import "testing"

func TestNumUniqueEmails(t *testing.T) {
	want := 2
	got := numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
