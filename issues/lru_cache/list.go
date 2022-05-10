package lru_cache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	listItem := ListItem{
		Value: v,
		Next:  l.Front(),
		Prev:  nil,
	}

	if l.Front() != nil {
		l.Front().Prev = &listItem
	}

	l.front = &listItem

	if l.Back() == nil {
		l.back = &listItem
	}

	l.incLen()
	return &listItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	listItem := ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.Back(),
	}

	if l.Back() != nil {
		l.Back().Next = &listItem
	}

	l.back = &listItem

	if l.Front() == nil {
		l.front = &listItem
	}

	l.incLen()
	return &listItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	l.decrLen()
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) incLen() {
	l.len++
}

func (l *list) decrLen() {
	l.len--
}

func NewList() List {
	return new(list)
}
