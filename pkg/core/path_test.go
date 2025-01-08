package core

import "testing"

func TestRandShortLinkPath(t *testing.T) {
	const l = 6
	a, b := RandShortLinkPath(l), RandShortLinkPath(l)

	if len(a) != l {
		t.Error("Expected", l, "characters, got", len(a))
	} else if len(b) != l {
		t.Error("Expected", l, "characters, got", len(b))
	} else if a == b {
		t.Error("Expected 2 random strings to be different, got", a)
	}
}
