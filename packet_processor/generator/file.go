package generator

import (
	"fmt"
	"os"
)

func WriteRandomPairsToFile(fileName string, nSamples int) {
	samples := GenerateRandomPairsBase64(nSamples)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't create file: ", err)
	}
	defer file.Close()

	for _, pair := range samples {
		file.WriteString(fmt.Sprintf("%d , %s\n", pair.Id, pair.Text))

	}
}
