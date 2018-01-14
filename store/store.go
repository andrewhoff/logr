package store

import (
	"fmt"

	"github.com/andrewhoff/logr/config"
)

// // Item - An abstraction of an item that can be stored in
// type Item interface {
// 	String() string
// }

// // DataStore ...
// type DataStore interface {
// 	Create() Item
// 	Delete() Item
// }

// Store ... map of slices
type Store struct {
	ds map[int]*ItemList
}

// InternalDataStore ...
var InternalDataStore *Store

// Init ...
func Init() {
	InternalDataStore = &Store{
		make(map[int]*ItemList, config.NumPriorities),
	}
}

// Dequeue .. removes the next in line, highest priority item from the list
func (store *Store) Dequeue() *Item {
	if store.LenWithPriority(config.HighPriority) != 0 {
		return store.ds[config.HighPriority].Pop()
	}

	if store.LenWithPriority(config.MedPriority) != 0 {
		return store.ds[config.MedPriority].Pop()
	}

	if store.LenWithPriority(config.LowPriority) != 0 {
		return store.ds[config.LowPriority].Pop()
	}

	return nil
}

// Enqueue .. removes the next in line, highest priority item from the list
func (store *Store) Enqueue(item *Item) (*Item, error) {
	cap := config.Capacity()

	if store.Len() == cap {
		return nil, fmt.Errorf("Store has reached capacity of %d", cap)
	}

	store.ds[item.Priority].Add(item)

	return item, nil
}

// Empty ...
func (store *Store) Empty() bool {
	return store.Len() == 0
}

// Len .. how many items are in logging system, of all priorities
func (store *Store) Len() int {
	length := 0

	if _, ok := store.ds[config.LowPriority]; !ok {
		store.ds[config.LowPriority] = &ItemList{make([]*Item, 0)}
	}

	if _, ok := store.ds[config.MedPriority]; !ok {
		store.ds[config.MedPriority] = &ItemList{make([]*Item, 0)}
	}

	if _, ok := store.ds[config.HighPriority]; !ok {
		store.ds[config.HighPriority] = &ItemList{make([]*Item, 0)}
	}

	length += len(store.ds[config.LowPriority].Items)
	length += len(store.ds[config.MedPriority].Items)
	length += len(store.ds[config.HighPriority].Items)

	return length
}

// LenWithPriority .. how many items are in logging system, of specified priority
func (store *Store) LenWithPriority(priority int) int {
	return len(store.ds[priority].Items)
}

func (store *Store) String() string {
	str := ""

	for _, v := range store.ds {
		fmt.Println(v)
	}

	return str
}
