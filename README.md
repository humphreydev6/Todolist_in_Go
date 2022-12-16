# Todolist_in_Go

To build a todo list in Go, you will need to design a data structure to store the todo items and a set of functions to manipulate the list. Here is an example of how you could implement a simple todo list using a slice of strings to store the todo items:

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


This example defines a TodoList struct that has a slice of strings to store the todo items. It also has four methods: Add, Remove, MarkDone, and Print. The Add method adds a new item to the list, the Remove method removes an item from the list, the MarkDone method marks an item as done by appending "(done)" to the item, and the Print method prints the items in the list.

You can modify this example to fit your specific needs. For example, you could add a field to the TodoList struct to store the date that the item was added, or you could implement a method to sort the items in the list by priority.
