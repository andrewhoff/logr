package writers

import (
	"testing"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
)

func TestHighPriorityLog(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewHighPriorityLogWriter()
	writer.Log("hi")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) != 1 {
		t.Fail()
	}
}

func TestHighPriorityLogSystemFullWithOverwrites(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{
		Capacity:  1,
		Overwrite: true,
	})

	writer := NewHighPriorityLogWriter()
	writer.Log("hi")
	writer.Log("hello")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) != 1 {
		t.Fail()
	}

	dequeued := store.InternalDataStore.Dequeue()
	if dequeued.Value != "hello" {
		t.Fail()
	}
}

func TestHighPriorityLogSystemFullWithoutOverwrites(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{
		Capacity:  1,
		Overwrite: false,
	})

	writer := NewHighPriorityLogWriter()
	writer.Log("hi")
	writer.Log("hello")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) != 1 {
		t.Fail()
	}

	dequeued := store.InternalDataStore.Dequeue()
	if dequeued.Value != "hi" {
		t.Fail()
	}
}
