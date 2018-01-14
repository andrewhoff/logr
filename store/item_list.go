package store

import "fmt"

// ItemList ....
type ItemList struct {
	Items []*Item
}

// Item ...
type Item struct {
	Priority int
	Value    string
}

// Add ...
func (list *ItemList) Add(item *Item) {
	list.Items = append(list.Items, item)
}

// Pop ...
func (list *ItemList) Pop() *Item {
	item := list.Items[0]
	list.Items = list.Items[1:]

	return item
}

func (list *ItemList) String() string {
	str := ""

	for _, v := range list.Items {
		str += fmt.Sprintf("%+v", v)
	}

	return str
}
