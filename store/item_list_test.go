package store

import (
	"testing"
	"time"

	"github.com/andrewhoff/logr/config"
)

func TestListAdd(t *testing.T) {
	list := NewItemList()

	item := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	list.Add(item)

	if list.Len() != 1 {
		t.Fail()
	}
}

func TestListPop(t *testing.T) {
	list := NewItemList()

	item := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	list.Add(item)

	if list.Len() != 1 {
		t.Fail()
	}

	popped := list.Pop()

	if popped != item {
		t.Fail()
	}

	if list.Len() != 0 {
		t.Fail()
	}
}
