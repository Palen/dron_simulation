package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/Palen/drone_simulation/pkg/config"
	"github.com/Palen/drone_simulation/pkg/dispatcher"
	"github.com/Palen/drone_simulation/pkg/geo"
	"github.com/Palen/drone_simulation/pkg/producers"
	"github.com/Palen/drone_simulation/pkg/subscribers"
)

func main() {
	// Parse config
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Instanttine dispatcher
	dispatcherChannel := make(dispatcher.DispatcherChan)

	// Read checkpoints file
	checkpts := geo.NewCheckPointsFromFile(conf.CheckPointFile)

	// Read subscribers files dir
	subscribersFiles, err := ioutil.ReadDir(conf.SubscribersDir)
	if err != nil {
		log.Fatal(err)
	}
	// Create subscribers slice
	subs := make(subscribers.Subscribers)
	// len(waiter) = len(subscribers)
	var waiter sync.WaitGroup
	for _, file := range subscribersFiles {
		idStr := strings.Split(file.Name(), ".csv")[0]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			log.Fatal("Invalid subscriber file name")
		}
		filePath := filepath.Join(conf.SubscribersDir, file.Name())
		fileReader := producers.NewFileReader(filePath, &dispatcherChannel)
		sub := subscribers.NewDrone(checkpts, conf.Drone.MaxSize, id, conf.Drone.Speed,
			conf.Drone.Perimeter)
		subs[id] = sub
		go sub.Subscribe()
		go fileReader.Read()
		waiter.Add(1)
	}
	// Start dispatching
	dispatcher := dispatcher.New(&dispatcherChannel)
	go dispatcher.Start(subs, &waiter)
	waiter.Wait()

}
