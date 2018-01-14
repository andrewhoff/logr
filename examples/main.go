package main

import (
	"fmt"

	"github.com/andrewhoff/logr"
	"github.com/andrewhoff/logr/config"
)

func main() {
	logr.InitWithOpts(&config.Opts{})

	reader, err := logr.NewLogReader()
	if err != nil {
		panic(err)
	}

	writer, err := logr.NewLogWriter()
	if err != nil {
		panic(err)
	}

	writer.Log(logr.HighPriority, "hello high priority")
	writer.Log(logr.HighPriority, "hello high prioritdadfasdfy")
	writer.Log(logr.MedPriority, "hello high prioriasdfaskhbvasldvsty")

	fmt.Println(reader.Get())
	fmt.Println(reader.Get())
	fmt.Println(reader.Get())
}
