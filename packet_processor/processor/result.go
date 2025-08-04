package processor

import (
	"strconv"
	"strings"
)

type Result struct {
	id           int
	count        int
	persianCount int
	englishCount int
	englishRatio int
}

func (r *Result) GetCsvString() string {
	s := []string{
		strconv.Itoa(r.id),
		strconv.Itoa(r.count),
		strconv.Itoa(r.persianCount),
		strconv.Itoa(r.englishCount),
		strconv.Itoa(r.englishRatio),
	}
	return strings.Join(s, ",")
}
