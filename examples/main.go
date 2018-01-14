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

	writer.Log(logr.HighPriority, "hello super high priority")
	writer.Log(logr.HighPriority, "hello 1")
	writer.Log(logr.HighPriority, "hello 2")
	writer.Log(logr.HighPriority, "hello 3")
	writer.Log(logr.MedPriority, "hello mid priority")
	writer.Log(logr.MedPriority, "hello mid priority")
	writer.Log(logr.HighPriority, "hello 4")
	writer.Log(logr.MedPriority, "hello mid priority")
	writer.Log(logr.LowPriority, "hello low priority")
	writer.Log(logr.MedPriority, "hello mid priority")

	gotten := reader.Get()
	fmt.Println(gotten)
	for gotten != "" {
		fmt.Println(gotten)
		gotten = reader.Get()
	}
}
