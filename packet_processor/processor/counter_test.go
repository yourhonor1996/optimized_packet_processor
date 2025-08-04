package processor

import (
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/shared"
	"os"
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

	outputFile, _ := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	c := NewCounter(shared.FileName)
	defer c.Close()
	for {
		result, err := c.CountNextPair()
		if err != nil {
			fmt.Println("file finished")
			break
		}
		_, err2 := outputFile.WriteString(result.GetCsvString() + "\n")
		if err2 != nil {
			fmt.Println("error writing to file")
		}
	}
}
