package main

import (
	"fmt"
	"strings"
)

type TodoList struct {
	items []string
}

func (l *TodoList) Add(item string) {
	l.items = append(l.items, item)
}

func (l *TodoList) Remove(item string) {
	for i, v := range l.items {
		if v == item {
			l.items = append(l.items[:i], l.items[i+1:]...)
			return
		}
	}
}

func (l *TodoList) MarkDone(item string) {
	for i, v := range l.items {
		if v == item {
			l.items[i] = strings.Replace(v, item, item+" (done)", 1)
			return
		}
	}
}

func (l *TodoList) Print() {
	for _, v := range l.items {
		fmt.Println(v)
	}
}

func main() {
	list := &TodoList{}
	list.Add("item1")
	list.Add("item2")
	list.Add("item3")
	list.Print()
	list.MarkDone("item2")
	list.Print()
	list.Remove("item3")
	list.Print()
}
