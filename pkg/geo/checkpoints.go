package geo

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// CheckPoint respresents an object with location
type CheckPoint struct {
	Name  string
	Coord *Coord
}

func NewCheckPointsFromFile(fileStr string) []*CheckPoint {
	var checkpoints []*CheckPoint
	file, err := os.Open(fileStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File, fileStr string) {
		err := file.Close()
		if err != nil {
			log.Printf("Error while closing file %s with err: %s", fileStr, err)
		}
	}(file, fileStr)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		coord, err := LatLonToCoords(fields[1], fields[2])
		if err != nil {
			log.Printf("Error pasring checkpoints file in line: %s with err: %s", line, err)
			continue
		}
		checkpoints = append(checkpoints, &CheckPoint{Name: fields[0], Coord: coord})
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error closing scanner for file: %s with err: %s", fileStr, err)
	}
	return checkpoints
}
