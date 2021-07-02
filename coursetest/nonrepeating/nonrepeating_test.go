package mian

import "testing"

func TestNonRepeating(t *testing.T) {
	tests := []struct {
		s string
		a int
	}{
		{"abcabc", 3},
		{"abcdefa", 6},
	}

	for _, tt := range tests {
		if acutal := MaxSubStringNoRepeat(tt.s); acutal != tt.a {
			t.Errorf("MaxSubStringNoRepeat(%s); excepted %d, but got %d", tt.s, acutal, tt.a)
		}
	}
}
