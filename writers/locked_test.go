package writers

import (
	"testing"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
)

func TestLockedLogHighPri(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewLockedPriorityLogWriter(config.HighPriority)
	writer.Log("hi")
	writer.Log("hello")
	writer.Log("hello again")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) != 3 {
		t.Fail()
	}
}

func TestLockedLogMedPri(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewLockedPriorityLogWriter(config.MedPriority)
	writer.Log("hi")
	writer.Log("hello")
	writer.Log("hello again")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) != 3 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) > 0 {
		t.Fail()
	}
}

func TestLockedLogLowPri(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewLockedPriorityLogWriter(config.LowPriority)
	writer.Log("hi")
	writer.Log("hello")
	writer.Log("hello again")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) != 3 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) > 0 {
		t.Fail()
	}
}

func TestLockedLogSystemFullWithOverwrites(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{
		Capacity:  1,
		Overwrite: true,
	})

	writer := NewLockedPriorityLogWriter(config.LowPriority)
	writer.Log("hi")
	writer.Log("hello")
	writer.Log("hello again")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) > 0 {
		t.Fail()
	}
}

func TestLockedLogSystemFullWithoutOverwrites(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{
		Capacity:  1,
		Overwrite: false,
	})

	writer := NewLockedPriorityLogWriter(config.LowPriority)
	writer.Log("hi")
	writer.Log("hello")
	writer.Log("hello again")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.MedPriority) > 0 {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.HighPriority) > 0 {
		t.Fail()
	}
}
