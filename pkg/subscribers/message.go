package subscribers

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Palen/drone_simulation/pkg/geo"
)

type MessageChannel chan *Message

type Message struct {
	Id    uint64
	Time  time.Time
	Coord geo.Coord
}

func NewMessage(line string) (*Message, error) {
	fields := strings.Split(line, ",")
	if len(fields) == 4 {
		id, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, errors.New("id format not valid")
		}
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, strings.ReplaceAll(fields[3], `"`, ""))
		if err != nil {
			return nil, errors.New("time format not valid")
		}
		coords, err := geo.LatLonToCoords(strings.ReplaceAll(fields[1], `"`, ""),
			strings.ReplaceAll(fields[2], `"`, ""))
		if err != nil {
			return nil, errors.New("coordinates not valid")
		}
		message := Message{Id: id, Coord: *coords, Time: t}
		return &message, nil
	} else {
		return nil, errors.New("line format not valid")
	}
}
