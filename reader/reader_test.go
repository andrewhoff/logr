package reader

import (
	"testing"
	"time"

	"github.com/andrewhoff/logr/config"
	"github.com/andrewhoff/logr/store"
	"github.com/fatih/color"
)

func TestGet(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	item := &store.Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	_, err := store.InternalDataStore.Enqueue(item)
	if err != nil {
		t.Error(err)
	}

	reader := NewLogReader()
	msg := reader.Get()
	if msg != color.GreenString("[LOW PRIORITY] - %s\n", item.Value) {
		t.Fail()
	}
}

func TestGetFromEmptyLogs(t *testing.T) {
	store.Init()
	config.SetOpts(&config.Opts{})

	reader := NewLogReader()
	msg := reader.Get()
	if msg != "" {
		t.Fail()
	}
}
