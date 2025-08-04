package generator

import (
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"
	"math/rand"
)

func randomFromRanges() rune {
	r := unicodeRanges[rand.Intn(len(unicodeRanges))]
	return r.start + rune(rand.Intn(int(r.end-r.start)+1))
}

func generateString(max_length int) string {
	randomLength := 1 + rand.Intn(max_length)
	id := make([]rune, randomLength)
	for i := range id {
		id[i] = randomFromRanges()
	}
	v := string(id)
	if len(v) > max_length {
		return v[:max_length]
	}
	return v
}

func GenerateEncodedBase64(length int) string {
	return utils.EncodeBase64(generateString(length))
}

func GenerateId(start int, stop int) int {
	return start + rand.Intn(stop-start-1)
}
