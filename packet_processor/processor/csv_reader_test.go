package processor

import (
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/shared"
	"testing"
)

func TestGetNextLine(t *testing.T) {
	fp := NewCsvReader(shared.FileName)
	pair, err := fp.GetNextPair()
	if err != nil {
		t.Errorf("Error reading even the first pair! :%v", err)
	}
	defer fp.Close()
	fmt.Println(pair)
}

func TestGetAllLines(t *testing.T) {
	fp := NewCsvReader(shared.FileName)
	defer fp.Close()
	for {
		pair, err := fp.GetNextPair()
		if err != nil {
			fmt.Println("finished file", err)
			break
		}
		fmt.Println(pair)
	}
	fmt.Println("Finished processing file.")
}
