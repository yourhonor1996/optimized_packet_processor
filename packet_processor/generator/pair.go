package generator

import "github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/shared"

func NewRandomPair(start int, end int) shared.Pair {
	return shared.Pair{
		Id:   generateId(start, end),
		Text: generateString(textLength),
	}
}

func GenerateRandomPairs(nSamples int) []shared.Pair {
	pairs := make([]shared.Pair, nSamples)
	for i := range pairs {
		pairs[i] = NewRandomPair(idRangeStart, idRangeEnd)
	}
	return pairs
}

func GenerateRandomPair(start int, end int) shared.Pair {
	return NewRandomPair(start, end)
}
