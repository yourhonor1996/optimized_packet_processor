package processor

import (
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"
	"os"
)

func CountAndWriteToFileOneByOne(inputFile, outputFile string) {
	output, _ := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	c := NewCounter(inputFile)
	defer c.Close()
	defer output.Close()
	for {
		result, err := c.CountNextPair()
		if err != nil {
			fmt.Println("file finished")
			break
		}
		output.WriteString(result.GetCsvString() + "\n")
		//if err2 != nil {
		//	fmt.Println("error writing to file")
		//}
	}
}

func CountAndWriteToFileInOneBatch(inputFile, outputFile string, nSamples int) {
	t := utils.Timer{}
	t.Start()

	fmt.Println("<Read All>")

	output, _ := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	c := NewCounter(inputFile)
	defer c.Close()
	defer output.Close()
	resultsStructs := make([]Result, nSamples)
	resultsStrings := make([]string, nSamples)

	for i := 0; i < nSamples; i++ {
		result, err := c.CountNextPair()
		if err != nil {
			fmt.Println("file finished")
			break
		}
		resultsStructs[i] = result
	}
	fmt.Printf("Time it took to Read All: <%d> Milliseconds.\n", t.DiffMilli())

	fmt.Println("- - - - - - - - - - - - - - - - - - - -")

	fmt.Println("<Process All>")
	t.Start()
	for i := 0; i < nSamples; i++ {
		resultsStrings[i] = resultsStructs[i].GetCsvString()
	}
	fmt.Printf("Time it took to Process All: <%d> Milliseconds.\n", t.DiffMilli())

	fmt.Println("- - - - - - - - - - - - - - - - - - - -")
	fmt.Println("<Write All>")
	t.Start()
	for _, result := range resultsStrings {
		_, e := output.WriteString(result + "\n")
		if e != nil {
			fmt.Println("Error writing to file")
		}
	}
	fmt.Printf("Time it took to Write All: <%d> Milliseconds.\n", t.DiffMilli())
}
