package subscribers

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Palen/drone_simulation/pkg/geo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var barnaCoord = geo.Coord{Lat: 41.385063, Lon: 2.173404}
var madCoord = geo.Coord{Lat: 40.416775, Lon: -3.703790}

func TestDroneSuite(t *testing.T) {
	suite.Run(t, new(DroneSuite))
}

type DroneSuite struct {
	suite.Suite
	d *Drone
}

func (s *DroneSuite) BeforeTest(suiteName, testName string) {
	id := uint64(99)
	//Checkpoints are Barcelona and Madrid
	checkpts := []*geo.CheckPoint{{Name: "Barecelona", Coord: &barnaCoord},
		{Name: "Madrid", Coord: &madCoord}}

	// Drone speed light
	speed := 300000000.00

	// permiter 15km
	perimeter := 15000.00
	s.d = NewDrone(checkpts, 1, id, speed, perimeter)
	go s.d.Subscribe()

}
func (s *DroneSuite) AfterTest(suiteName, testName string) {
	log.SetOutput(os.Stderr)
}

func newMessage(coord *geo.Coord) Message {
	return Message{Id: uint64(99), Time: time.Now(), Coord: coord}
}

func (s *DroneSuite) Test_All_Drone() {
	// route will be barcelona-terrassa-madrid
	messageBarna := newMessage(&barnaCoord)
	var terCoord = geo.Coord{Lat: 41.385063, Lon: 2.173404}
	messageTer := newMessage(&terCoord)
	messageMad := newMessage(&madCoord)
	var buf bytes.Buffer
	log.SetOutput(&buf)
	s.d.Send(&messageBarna)
	s.d.Send(&messageTer)
	s.d.Send(&messageMad)
	time.Sleep(1 * time.Second)
	logs := strings.Split(buf.String(), "\n")
	reports := 0
	for _, l := range logs {
		if strings.Contains(l, "Traffic Report") {
			reports++
		}
	}
	assert.Equal(s.T(), 2, reports)
	log.SetOutput(os.Stderr)
}
