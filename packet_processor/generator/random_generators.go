package generator

import (
	"math/rand"
)

func randomFromRanges() rune {
	r := unicodeRanges[rand.Intn(len(unicodeRanges))]
	return r.start + rune(rand.Intn(int(r.end-r.start)+1))
}

func generateString(length int) string {
	randomLength := 1 + rand.Intn(length)
	id := make([]rune, randomLength)
	for i := range id {
		id[i] = randomFromRanges()
	}
	v := string(id)
	if len(v) > length {
		return v[:length]
	}
	return v
}

func generateId(start int, stop int) int {
	return start + rand.Intn(stop-start-1)
}
