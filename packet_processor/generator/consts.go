package generator

var unicodeRanges = []struct {
	start, end rune
}{
	{0x0041, 0x005A}, // A-Z
	{0x0061, 0x007A}, // a-z
	{0x0600, 0x06FF}, // Arabic and Persian
}

const (
	textLength   = 70
	idRangeStart = 100000
	idRangeEnd   = 1000000
)

const (
	randomFilePrefix = "RandomSample"
)
