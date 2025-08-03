package generator

import (
	"fmt"
	"testing"
)

func TestGenerateSampleFileName(t *testing.T) {
	fileName := GenerateRandomSampleFileName()
	fmt.Println(fileName)
}

func TestGenerateAFileWithRandomSamples(t *testing.T) {
	fileName := GenerateRandomSampleFileName()
	var numberOfSamples uint32 = 1000
	WriteRandomPairsToFile(fileName, numberOfSamples)
	//TODO-CONTINUEFROMHERE continue from this test...
}
