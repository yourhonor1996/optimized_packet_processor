package processor

import (
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"
)

type Counter struct {
	reader CsvReader
}

func (c *Counter) Close() error {
	return c.reader.Close()
}

func (c *Counter) CountNextPair() (Result, error) {
	pair, err := c.reader.GetNextPairDecoded()
	if err != nil {
		return Result{}, err
	}

	englishCount := utils.CountEnglishChars(&pair.Text)
	length := len(pair.Text)
	res := Result{
		id:           pair.Id,
		count:        length,
		persianCount: length - englishCount,
		englishCount: englishCount,
		englishRatio: englishCount * 100 / length,
	}
	return res, nil
}

// TODO handle if the file doesn't exist
func NewCounter(name string) *Counter {
	return &Counter{reader: NewCsvReader(name)}
}
