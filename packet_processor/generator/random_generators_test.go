package generator

import (
	"fmt"
	"testing"
	"unicode/utf16"
)

func TestCompareUtf8AndUtf16(t *testing.T) {
	s := "Aش😊"
	bytes := []byte(s)
	runes := []rune(s)
	encode := utf16.Encode(runes)
	fmt.Println(bytes)  // UTF-8 bytes
	fmt.Println(runes)  // Unicode code points
	fmt.Println(encode) // UTF-16 units
	fmt.Println("sdf")
}

func TestGenerateRandomCharFromRanges(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := randomFromRanges()
		println(string(str))
	}
}

func TestGenerateRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := generateString(70)
		println(str)
	}
}

func TestGenerateId(t *testing.T) {
	for i := 0; i < 10; i++ {
		str := GenerateId(100000, 1000000)
		println(str)
	}
}

func TestGeneratePair(t *testing.T) {
	for i := 0; i < 20; i++ {
		pair := NewRandomPairBase64(100000, 1000000)
		fmt.Println(pair)
	}
}

func TestGeneratedStringHasRandomLength(t *testing.T) {
	for i := 0; i < 100; i++ {
		str := generateString(textLength)
		if len(str) > textLength || len(str) == 0 {
			t.Errorf("Generated string does not have proper length.")
		}
	}
}
