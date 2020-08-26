package checkpoints

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64
	Lon float64
}

// CheckPoint respresents an object with location
type CheckPoint struct {
	Name  string
	coord *Coord
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth mesarued in meters.
func (c *CheckPoint) Distance(p *Coord) float64 {
	lat1 := degreesToRadians(c.coord.Lat)
	lon1 := degreesToRadians(c.coord.Lon)
	lat2 := degreesToRadians(p.Lat)
	lon2 := degreesToRadians(p.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	d := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km := d * earthRaidusKm

	return km * 1000
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func LatLonToCoords(lat, lon string) (*Coord, error) {
	latitude, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return nil, err
	}
	longitude, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return nil, err
	}
	coord := Coord{Lat: latitude, Lon: longitude}
	return &coord, nil
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
		checkpoints = append(checkpoints, &CheckPoint{Name: fields[0], coord: coord})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return checkpoints
}
