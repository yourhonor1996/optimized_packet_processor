package main

import (
	"flag"
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/generator"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/processor"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"
	"path/filepath"
)

func main() {
	input := flag.String("i", "", "Path to input file - if you don't specify this it will be auto generated")
	output := flag.String("o", "", "Path to output file - if you don't specify this it will be auto generated")
	nSamples := flag.Int("n", 100, "Number of samples to generate/process - if you specify an input file, this many samples will be processed from that file. Default is 100")

	flag.Parse()

	t := utils.Timer{}

	if *input == "" {
		fmt.Println("======================================")
		*input = filepath.Join(".", generator.GenerateRandomSampleFileName())
		fmt.Println("Generating file, please wait ... ", *input)
		t.Start()
		generator.WriteRandomPairsToFile(*input, *nSamples)
		t.PrintDiffMilli("WriteRandomPairsToFile")

	}

	if *output == "" {
		*output = processor.GenerateResultFileName(*input)
		fmt.Println("Generate output file: ", *output)
	}

	fmt.Println("Using input file:", *input)
	fmt.Println("Using output file:", *output)
	fmt.Println("Number of samples that will be processes:", *nSamples)

	fmt.Println("======================================")
	fmt.Printf("Processing <%d> samples from file <%s>\n", *nSamples, *input)
	fmt.Println("======================================")

	fmt.Println("Read One -> Process One -> Write One")
	t.Start()
	processor.CountAndWriteToFileOneByOne(*input, *output)
	t.PrintDiffMilli("CountAndWriteToFileOneByOne")

	fmt.Println("======================================")

	fmt.Println("Read All -> Process All -> Write All")
	t.Start()
	processor.CountAndWriteToFileInOneBatch(*input, *output, *nSamples)
	t.PrintDiffMilli("CountAndWriteToFileInOneBatch")

	fmt.Println("======================================")
}
