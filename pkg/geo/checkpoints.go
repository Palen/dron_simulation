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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		coord, err := LatLonToCoords(fields[1], fields[2])
		if err != nil {
			log.Println("error pasring checkpoints file with err:", err)
			continue
		}
		checkpoints = append(checkpoints, &CheckPoint{Name: fields[0], Coord: coord})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return checkpoints
}
