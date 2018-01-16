package store

import (
	"fmt"
	"time"
)

// ItemList - list of log items
type ItemList struct {
	Items []*Item
}

// NewItemList ...
func NewItemList() *ItemList {
	return &ItemList{
		Items: make([]*Item, 0),
	}
}

// Item - A log item, kind of the 'atom' of the system
type Item struct {
	Priority int
	Value    string
	DateTime time.Time
}

// Add - append to end of list
func (list *ItemList) Add(item *Item) {
	list.Items = append(list.Items, item)
}

// Pop - remove from top of list (oldest)
func (list *ItemList) Pop() *Item {
	item := list.Items[0]
	list.Items = list.Items[1:]

	return item
}

// PopBack - remove from end of list (most recent)
func (list *ItemList) PopBack() *Item {
	item := list.Items[len(list.Items)-1]
	list.Items = list.Items[:len(list.Items)-1]

	return item
}

// Len ...
func (list *ItemList) Len() int {
	return len(list.Items)
}

func (list *ItemList) String() string {
	str := ""

	for _, v := range list.Items {
		str += fmt.Sprintf("%+v", v)
	}

	return str
}
