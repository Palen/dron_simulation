package producers

import (
	"bufio"
	"log"
	"os"

	"github.com/Palen/drone_simulation/pkg/dispatcher"
)

type FileReader struct {
	fileStr        string
	dispatcherChan *dispatcher.DispatcherChan
}

func (r *FileReader) Read() {
	file, err := os.Open(r.fileStr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*r.dispatcherChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func NewFileReader(fileStr string, channel *dispatcher.DispatcherChan) *FileReader {
	fileReader := FileReader{fileStr: fileStr, dispatcherChan: channel}
	return &fileReader

}
