package processor

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/shared"
	"github.com/yourhonor1996/optimized_packet_processor/packet_proecssor/utils"
	"os"
	"strconv"
	"strings"
)

type CsvReader struct {
	fileName   string
	file       *os.File
	reader     *bufio.Reader
	lineNumber int
}

func (p *CsvReader) Close() error {
	return p.file.Close()
}

func (p *CsvReader) GetNextPairDecoded() (shared.Pair, error) {
	s, err := p.reader.ReadString('\n')
	if err != nil {
		defer p.file.Close()
		return shared.Pair{}, errors.New("file is finished")
	}

	split := strings.Split(s, ",")
	if len(split) != 2 {
		fmt.Printf("The format of the file is not csv compatible at line <%d>: %s", p.lineNumber, s)
	}
	text := strings.TrimSpace(split[1])
	idInt, err := strconv.Atoi(strings.TrimSpace(split[0][:len(split[0])-1]))
	if err != nil {
		return shared.Pair{}, errors.New(fmt.Sprintf("The id is invalid at line <%d>", p.lineNumber))
	}

	p.lineNumber++

	return shared.Pair{Id: idInt, Text: utils.DecodeBase64(text)}, nil

}

func NewCsvReader(name string) CsvReader {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("Error opening file! :%v", err)
	}
	return CsvReader{reader: bufio.NewReader(file), file: file, fileName: name, lineNumber: 1}
}
