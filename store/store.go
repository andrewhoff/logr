package store

import (
	"fmt"
	"sync"

	"github.com/andrewhoff/logr/config"
)

// Store ... The highest level abstraction for where we put logged messages
type Store struct {
	ds    map[int]*ItemList
	mutex *sync.Mutex
}

// InternalDataStore ...
var InternalDataStore *Store

// Init ...
func Init() {
	InternalDataStore = &Store{
		make(map[int]*ItemList, config.NumPriorities),
		&sync.Mutex{},
	}
}

// Enqueue .. Adds the item to the store
func (store *Store) Enqueue(item *Item) (*Item, error) {
	cap := config.Capacity()

	if store.Len() == cap {
		if config.ShouldOverwrite() {
			store.DequeueOldestLowest()
			return store.Enqueue(item)
		}

		return nil, fmt.Errorf("Store has reached capacity of %d", cap)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.ds[item.Priority].Add(item)

	return item, nil
}

// Dequeue .. removes the most recent highest priority item from the list
func (store *Store) Dequeue() *Item {
	if store.LenWithPriority(config.HighPriority) != 0 {
		store.mutex.Lock()
		defer store.mutex.Unlock()
		return store.ds[config.HighPriority].PopBack()
	}

	if store.LenWithPriority(config.MedPriority) != 0 {
		store.mutex.Lock()
		defer store.mutex.Unlock()
		return store.ds[config.MedPriority].PopBack()
	}

	if store.LenWithPriority(config.LowPriority) != 0 {
		store.mutex.Lock()
		defer store.mutex.Unlock()
		return store.ds[config.LowPriority].PopBack()
	}

	return nil
}

// DequeueOldestLowest .. removes the oldest, lowest priority message possible, to make space for when the system is at capacity
func (store *Store) DequeueOldestLowest() *Item {
	if store.LenWithPriority(config.LowPriority) != 0 {
		store.mutex.Lock()
		defer store.mutex.Unlock()
		return store.ds[config.LowPriority].Pop()
	}

	if store.LenWithPriority(config.MedPriority) != 0 {
		store.mutex.Lock()
		defer store.mutex.Unlock()
		return store.ds[config.MedPriority].Pop()
	}

	if store.LenWithPriority(config.HighPriority) != 0 {
		store.mutex.Lock()
		defer store.mutex.Unlock()
		return store.ds[config.HighPriority].Pop()
	}

	return nil
}

// Empty ...
func (store *Store) Empty() bool {
	return store.Len() == 0
}

// Len .. how many items are in logging system, of all priorities
func (store *Store) Len() int {
	store.mutex.Lock()
	defer store.mutex.Unlock()

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
	store.mutex.Lock()
	defer store.mutex.Unlock()

	return len(store.ds[priority].Items)
}

func (store *Store) String() string {
	str := ""

	for _, v := range store.ds {
		fmt.Println(v)
	}

	return str
}
