package generator

import "github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/shared"

func NewRandomPairBase64(start int, end int) shared.Pair {
	return shared.Pair{
		Id:   GenerateId(start, end),
		Text: GenerateEncodedBase64(textLength),
	}
}

func GenerateRandomPairsBase64(nSamples int) []shared.Pair {
	pairs := make([]shared.Pair, nSamples)
	for i := range pairs {
		pairs[i] = NewRandomPairBase64(idRangeStart, idRangeEnd)
	}
	return pairs
}
