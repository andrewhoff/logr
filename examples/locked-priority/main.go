package main

import (
	"fmt"

	"github.com/andrewhoff/logr"
	"github.com/andrewhoff/logr/config"
)

func main() {
	logr.InitWithOpts(&config.Opts{
		Overwrite: true,
	})

	reader, err := logr.NewLogReader()
	if err != nil {
		panic(err)
	}

	lowPriorityWriter, err := logr.NewLockedPriorityLogWriter(config.LowPriority)
	if err != nil {
		panic(err)
	}

	highPriorityWriter, err := logr.NewLockedPriorityLogWriter(config.HighPriority)
	if err != nil {
		panic(err)
	}

	lowPriorityWriter.Log("This is a lowish priority message")
	lowPriorityWriter.Log("This is a low priority message")
	lowPriorityWriter.Log("This is a very low priority message")

	highPriorityWriter.Log("This is a very high priority message")
	highPriorityWriter.Log("This is a high - 1 low priority message")
	highPriorityWriter.Log("This is a high - 2 low priority message")
	highPriorityWriter.Log("This is a high - 3 priority message")
	highPriorityWriter.Log("This is a high-ish priority message")

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
