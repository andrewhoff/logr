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

	lazyLogWriter, err := logr.NewLazyLogWriter(config.LowPriority)
	if err != nil {
		panic(err)
	}

	lazyLogWriter.Log("abc")
	reader.Get() // this will return empty
	lazyLogWriter.Log("def")
	lazyLogWriter.Log("ghi")

	if err := lazyLogWriter.Flush(); err != nil {
		panic(err)
	}

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
