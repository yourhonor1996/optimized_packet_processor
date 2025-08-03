package generator

type Pair struct {
	id   uint32
	text string
}

func NewRandomPair(start int, end int) Pair {
	return Pair{
		id:   generateId(start, end),
		text: generateString(textLength),
	}
}

func CreateRandomPairs(nSamples uint32) []Pair {
	pairs := make([]Pair, nSamples)
	for i := range pairs {
		pairs[i] = NewRandomPair(idRangeStart, idRangeEnd)
	}
	return pairs
}

func GeneratePair(start int, end int) Pair {
	return NewRandomPair(start, end)
}
