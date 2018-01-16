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

func TestListPopBackOneItem(t *testing.T) {
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

	popped := list.PopBack()

	if popped != item {
		t.Fail()
	}

	if list.Len() != 0 {
		t.Fail()
	}
}

func TestListPopBackMoreThanOneItem(t *testing.T) {
	list := NewItemList()

	item1 := &Item{
		Priority: config.LowPriority,
		Value:    "hi",
		DateTime: time.Now(),
	}

	item2 := &Item{
		Priority: config.LowPriority,
		Value:    "hello",
		DateTime: time.Now(),
	}

	list.Add(item1)
	list.Add(item2)

	if list.Len() != 2 {
		t.Fail()
	}

	popped := list.PopBack()

	if popped != item2 {
		t.Fail()
	}

	if list.Len() != 1 {
		t.Fail()
	}
}
