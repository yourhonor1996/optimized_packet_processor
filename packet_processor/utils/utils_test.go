package utils

import (
	"fmt"
	"testing"
)

func TestGetFileNameBasedOnDateTime(t *testing.T) {
	fmt.Println(GetNowAsFormattedString())
}

func TestCountEnglishChars(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{"<UNK>", 3},
		{"<UNKsdf>", 6},
		{"<UNKsdfسیبمنت>", 6},
		{"<UNKsdfسیبمنت؟>", 6},
	}

	for _, test := range tests {
		result := CountEnglishChars(&test.input)
		if result != test.exp {
			t.Errorf("TestCountEnglishChars failed: input: %s, exp: %d, got: %d",
				test.input, test.exp, result,
			)
		}
	}
}
