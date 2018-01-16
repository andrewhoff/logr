package writers

import (
	"testing"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
)

func TestLogDoesntSendLog(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewLazyLogWriter(config.HighPriority)
	writer.Log("a")
	writer.Log("b")
	writer.Log("c")

	if !store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) > 0 {
		t.Fail()
	}
}

func TestFlushSendsLog(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewLazyLogWriter(config.HighPriority)
	writer.Log("a")
	writer.Log("b")
	writer.Log("c")

	if !store.InternalDataStore.Empty() {
		t.Fail()
	}

	if err := writer.Flush(); err != nil {
		t.Error(err)
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) == 0 {
		t.Fail()
	}

	item := store.InternalDataStore.Dequeue()
	if item.Value != "abc" {
		t.Fail()
	}

	if !store.InternalDataStore.Empty() {
		t.Fail()
	}
}

func TestReset(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewLazyLogWriter(config.HighPriority)
	writer.Log("a")
	writer.Log("b")
	writer.Log("c")

	if !store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) > 0 {
		t.Fail()
	}

	if err := writer.Flush(); err != nil {
		t.Error(err)
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) == 0 {
		t.Fail()
	}

	item := store.InternalDataStore.Dequeue()
	if item.Value != "abc" {
		t.Fail()
	}

	if !store.InternalDataStore.Empty() {
		t.Fail()
	}

	writer.Log("a")
	if err := writer.Flush(); err != nil {
		t.Error(err)
	}

	dequeued := store.InternalDataStore.Dequeue()
	if dequeued.Value != "a" {
		t.Fail()
	}
}
