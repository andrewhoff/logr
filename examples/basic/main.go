package main

import (
	"fmt"
	"time"

	"github.com/andrewhoff/logr"
	"github.com/andrewhoff/logr/config"
)

func main() {
	logr.InitWithOpts(&config.Opts{
		Capacity:  24,
		Overwrite: true,
	})

	reader, err := logr.NewLogReader()
	if err != nil {
		panic(err)
	}

	writer, err := logr.NewGenericLogWriter()
	if err != nil {
		panic(err)
	}

	go func() {
		if err := writer.Log(logr.HighPriority, "hello super high priority"); err != nil {
			fmt.Printf("Error encountered trying to write to logs: %v\n", err)
		}
	}()

	go func() {
		if err := writer.Log(logr.LowPriority, "hello low priority 1"); err != nil {
			fmt.Printf("Error encountered trying to write to logs: %v\n", err)
		}
	}()

	go func() {
		if err := writer.Log(logr.MedPriority, "hello mid priority 1"); err != nil {
			fmt.Printf("Error encountered trying to write to logs: %v\n", err)
		}
	}()

	go func() {
		if err := writer.Log(logr.HighPriority, "hello 1"); err != nil {
			fmt.Printf("Error encountered trying to write to logs: %v\n", err)
		}
	}()

	go func() {
		if err := writer.Log(logr.HighPriority, "hello 2"); err != nil {
			fmt.Printf("Error encountered trying to write to logs: %v\n", err)
		}
	}()

	go func() {
		if err := writer.Log(logr.HighPriority, "hello 3"); err != nil {
			fmt.Printf("Error encountered trying to write to logs: %v\n", err)
		}
	}()

	if err := writer.Log(logr.MedPriority, "hello mid priority 2"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}
	if err := writer.Log(logr.MedPriority, "hello mid priority 3"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}
	if err := writer.Log(logr.HighPriority, "hello 4"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}
	if err := writer.Log(logr.MedPriority, "hello mid priority 4"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}
	if err := writer.Log(logr.LowPriority, "hello low priority 2"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}
	if err := writer.Log(logr.LowPriority, "hello lowest priority"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}
	if err := writer.Log(logr.MedPriority, "hello mid priority 5"); err != nil {
		fmt.Printf("Error encountered trying to write to logs: %v\n", err)
	}

	time.Sleep(250 * time.Millisecond)

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
