package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Palen/drone_simulation/pkg/checkpoints"
	"github.com/Palen/drone_simulation/pkg/dispatcher"
	"github.com/Palen/drone_simulation/pkg/producers"
	"github.com/Palen/drone_simulation/pkg/subscribers"
)

func main() {
	// Parse flags
	subscribersFileDirPtr := flag.String("subscribersdir", "./data/subscribers/", "Subscribers files location")
	checkpointsFilePtr := flag.String("checkpointsfile", "./data/tube.csv", "Checkpoints file")
	flag.Parse()

	// Instanttine dispatcher
	dispatcherChannel := make(dispatcher.DispatcherChan)

	// Read checkpoints file
	checkpts := checkpoints.NewCheckPointsFromFile(*checkpointsFilePtr)

	// Read subscribers files dir
	subscribersFiles, err := ioutil.ReadDir(*subscribersFileDirPtr)
	if err != nil {
		log.Fatal(err)
	}

	subs := make(subscribers.Subscribers)
	for _, file := range subscribersFiles {
		idStr := strings.Split(file.Name(), ".csv")[0]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			log.Fatal("Invalid subscriber file name")
		}
		filePath := filepath.Join(*subscribersFileDirPtr, file.Name())
		fileReader := producers.NewFileReader(filePath, &dispatcherChannel)
		drone := subscribers.NewDrone(checkpts, 10, id)
		subs[id] = drone
		go drone.Subscribe()
		go fileReader.Read()
	}
	dispatcher := dispatcher.New(&dispatcherChannel)
	dispatcher.Start(subs)

}
