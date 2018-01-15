package writers

import (
	"testing"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
)

func TestGenericLog(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	writer := NewGenericLogWriter()
	writer.Log(config.LowPriority, "hi")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}
}

func TestGenericLogSystemFull(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{
		Capacity: 1,
	})

	writer := NewGenericLogWriter()
	writer.Log(config.LowPriority, "hi")
	writer.Log(config.LowPriority, "hi again")

	if store.InternalDataStore.Empty() {
		t.Fail()
	}

	if store.InternalDataStore.LenWithPriority(config.LowPriority) != 1 {
		t.Fail()
	}
}
