package producers

import (
	"bufio"
	"log"
	"os"

	"github.com/Palen/drone_simulation/pkg/dispatcher"
)

type FileReader struct {
	fileStr        string
	dispatcherChan dispatcher.DispatcherChan
}

func (r *FileReader) Read() {
	file, err := os.Open(r.fileStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File, fileStr string) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing %s with err: %s", fileStr, err)
		}
	}(file, r.fileStr)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r.dispatcherChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func NewFileReader(fileStr string, channel dispatcher.DispatcherChan) *FileReader {
	fileReader := FileReader{fileStr: fileStr, dispatcherChan: channel}
	return &fileReader

}
