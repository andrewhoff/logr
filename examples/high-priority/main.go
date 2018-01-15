package main

import (
	"fmt"

	"github.com/andrewhoff/logr"
	"github.com/andrewhoff/logr/config"
)

func main() {
	logr.InitWithOpts(&config.Opts{
		Capacity:  3,
		Overwrite: true,
	})

	reader, err := logr.NewLogReader()
	if err != nil {
		panic(err)
	}

	highPriWriter, err := logr.NewHighPriorityLogWriter()
	if err != nil {
		panic(err)
	}

	highPriWriter.Log("This is a very high priority message!")
	highPriWriter.Log("This is a very high priority message ... again")
	highPriWriter.Log("This is a very high priority message ... again again")

	gotten := reader.Get()
	if gotten != "" {
		fmt.Printf("Got log msg: %s", gotten)
	}
	for gotten != "" {
		gotten = reader.Get()
		if gotten != "" {
			fmt.Printf("Got log msg: %s", gotten)
		}
	}
}
