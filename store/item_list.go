package store

import (
	"fmt"
	"time"
)

// ItemList - list of log items
type ItemList struct {
	Items []*Item
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

// Pop - remove from top of list
func (list *ItemList) Pop() *Item {
	item := list.Items[0]
	list.Items = list.Items[1:]

	return item
}

// PopBack - pop from the back of the list
func (list *ItemList) PopBack() *Item {
	item := list.Items[len(list.Items)-1]
	list.Items = list.Items[:len(list.Items)-1]

	return item
}

func (list *ItemList) String() string {
	str := ""

	for _, v := range list.Items {
		str += fmt.Sprintf("%+v", v)
	}

	return str
}
