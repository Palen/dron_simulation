package subscribers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Palen/drone_simulation/pkg/geo"
)

type MessageChannel chan *Message

type Message struct {
	Id    uint64
	Time  time.Time
	Coord *geo.Coord
}

func NewMessage(line string) (*Message, error) {
	fields := strings.Split(line, ",")
	if len(fields) == 4 {
		id, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing id in message with err: %s", err)
		}
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, strings.ReplaceAll(fields[3], `"`, ""))
		if err != nil {
			return nil, fmt.Errorf("error parsing time in message with err: %s", err)
		}
		coords, err := geo.LatLonToCoords(strings.ReplaceAll(fields[1], `"`, ""),
			strings.ReplaceAll(fields[2], `"`, ""))
		if err != nil {
			return nil, err
		}
		message := Message{Id: id, Coord: coords, Time: t}
		return &message, nil
	} else {
		return nil, errors.New(fmt.Sprintf("line format not valid in line: %s", line))
	}
}
