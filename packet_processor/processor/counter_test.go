package processor

import (
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/shared"
	"path/filepath"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	c := NewCounter(shared.FileName)
	defer c.Close()
	for {
		result, err := c.CountNextPair()
		if err != nil {
			fmt.Println("file finished")
			break
		}
		fmt.Println(result)
	}
}

func TestCountAndWriteToFileOneByOne(t *testing.T) {
	directory := filepath.Dir(shared.FileName)
	newPath := filepath.Join(directory, "result_"+filepath.Base(shared.FileName))

	CountAndWriteToFileOneByOne(shared.FileName, newPath)
}
