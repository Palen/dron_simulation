package geo

import (
	"fmt"
	"math"
	"strconv"
)

const (
	earthRaidusM = 6371000 // radius of the earth in meters.
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64
	Lon float64
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth mesarued in meters.
func (c *Coord) Distance(p *Coord) float64 {
	lat1 := degreesToRadians(c.Lat)
	lon1 := degreesToRadians(c.Lon)
	lat2 := degreesToRadians(p.Lat)
	lon2 := degreesToRadians(p.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	d := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	m := d * earthRaidusM

	return m
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// creates new coord pointer from lat lon string
func LatLonToCoords(lat, lon string) (*Coord, error) {
	latitude, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return nil, fmt.Errorf("latitude fails: %s", err)
	}
	longitude, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return nil, fmt.Errorf("longitude fails: %s", err)
	}
	coord := Coord{Lat: latitude, Lon: longitude}
	return &coord, nil
}
