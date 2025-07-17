package words

import "testing"

func TestIsSynonymOfTwoLetterWord(t *testing.T) {
	testCases := []struct {
		word     string
		synonym  string
		expected bool
	}{
		{"yes", "si", true},
		{"ourselves", "us", true},
		{"we", "us", true},
		{"is", "be", true},
		{"are", "be", true},
		{"am", "be", true},
		{"hello", "", false},
	}

	for _, tc := range testCases {
		synonym, ok := IsSynonymOfTwoLetterWord(tc.word)
		if ok != tc.expected {
			t.Errorf("expected %t, but got %t for word %s", tc.expected, ok, tc.word)
		}
		if synonym != tc.synonym {
			t.Errorf("expected synonym %s, but got %s for word %s", tc.synonym, synonym, tc.word)
		}
	}
}
