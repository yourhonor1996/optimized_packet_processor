package generator

import (
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"
	"os"
)

func GenerateRandomSampleFileName() string {
	return randomFilePrefix + "_" + utils.GetNowAsFormattedString() + ".txt"
}

func WriteRandomPairsToFile(fileName string, nSamples uint32) {
	samples := CreateRandomPairs(nSamples)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't create file: ", err)
	}
	defer file.Close()

	for _, pair := range samples {
		file.WriteString(fmt.Sprintf("%d , %s\n", pair.id, pair.text))

	}
}
