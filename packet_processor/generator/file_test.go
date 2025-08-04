package generator

import (
	"fmt"
	"testing"
)

func TestGenerateSampleFileName(t *testing.T) {
	fileName := GenerateRandomSampleFileName()
	fmt.Println(fileName)
}

//func TestGenerateAFileWithRandomSamples(t *testing.T) {
//	fileName := GenerateRandomSampleFileName()
//	var numberOfSamples int = 100
//	WriteRandomPairsToFile(fileName, numberOfSamples)
//}
