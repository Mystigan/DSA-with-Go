package linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	Data interface{}
	Next *node
}

// linkedList is the actual linked list of nodes.
type linkedList struct {
	Length int
	Head   *node
}

// InitList initializes and returns the reference to a new linkedlist.
func InitList() *linkedList {
	return &linkedList{}
}

// Size returns the size/length of the Linked List.
func (ll *linkedList) Size() int {
	return ll.Length
}

// Add adds a node at the beginning of the Linked List
func (ll *linkedList) Add(data interface{}) bool {
	node := &node{
		Data: data,
	}
	if ll.Head == nil {
		ll.Head = node
		ll.Length++
		return true
	}
	node.Next = ll.Head
	ll.Head = node
	ll.Length++
	return true
}

// AddAt adds a node at given index in the Linked List. Returns an error if the index is invalid.
//Note: The list is zero indexed. Therefore the provided index must fall in the range 0 <= index <= SizeOfList.
func (ll *linkedList) AddAt(data interface{}, index int) error {
	node := &node{
		Data: data,
	}
	if ll.Head == nil {
		ll.Head = node
		return errors.New("invalid index. List is empty")
	}
	if index > ll.Size() || index < 0 {
		return errors.New("invalid index. index must be a number between 0 and list.Size()")
	}
	if index == 0 {
		node.Next = ll.Head
		ll.Head = node
		ll.Length++
		return nil
	}

	previousNode := ll.Head
	for index > 1 {
		previousNode = previousNode.Next
		index--
	}
	node.Next = previousNode.Next
	previousNode.Next = node
	ll.Length++
	return nil
}

// Find searches for the specified item in the list. If found, returns the index of the item, else returns an error.
func (ll *linkedList) Find(data interface{}) (int, error) {

	if ll.Head == nil {
		return 0, errors.New("item not found. List is empty")
	}
	index := 0
	currentItem := ll.Head
	if data == currentItem.Data {
		return 1, nil
	}
	for currentItem.Next != nil {
		currentItem = currentItem.Next
		index++
		if data == currentItem.Data {
			return index, nil
		}
	}
	return 0, errors.New("item not found")

}

// Print iterates through the list and prints each element. Returns an error if the list is empty.
func (ll *linkedList) Print() error {
	if ll.Head == nil {
		return errors.New("list is empty")
	}
	currentItem := ll.Head
	fmt.Printf("\n")
	fmt.Printf("%v ", currentItem.Data)
	for currentItem.Next != nil {
		currentItem = currentItem.Next
		fmt.Printf("%v ", currentItem.Data)
	}
	fmt.Printf("\n")
	return nil
}

// Remove deletes the first element(Head) from the list. Returns an error if the list is empty.
func (ll *linkedList) Remove() (interface{}, error) {
	if ll.Head == nil {
		return nil, errors.New("list is empty")
	}
	currentItem := ll.Head
	ll.Head = ll.Head.Next
	currentItem.Next = nil
	ll.Length--
	return currentItem.Data, nil

}

// RemoveFrom deletes the item present at specified index in the list. Return an error if the index is invalid.
// Note: List is zero indexed. Provided index must be in the range 0 <= index < SizeOfList.
func (ll *linkedList) RemoveFrom(index int) (interface{}, error) {
	if ll.Head == nil {
		return nil, errors.New("list is empty")
	}
	if index < 0 || index >= ll.Size() {
		return nil, errors.New("invalid index")
	}
	currentItem := ll.Head
	if index == 0 {
		ll.Head = currentItem.Next
		currentItem.Next = nil
		ll.Length--
		return currentItem.Data, nil
	}
	// Traverse to index in the list and remove that item.
	previousItem := ll.Head
	for index > 0 {
		currentItem, previousItem = currentItem.Next, currentItem
		index--
	}
	previousItem.Next = currentItem.Next
	currentItem.Next = nil
	ll.Length--
	return currentItem.Data, nil
}

// RemoveItem deletes the given item from the list. Returns an error if the list is empty or the item was not found.
func (ll *linkedList) RemoveItem(data interface{}) error {
	if ll.Head == nil {
		return errors.New("list is empty")
	}
	currentItem := ll.Head
	previousItem := ll.Head
	for currentItem.Next != nil {
		if currentItem.Data == data {
			previousItem.Next = currentItem.Next
			currentItem.Next = nil
			ll.Length--
			return nil
		}
		previousItem = currentItem
		currentItem = currentItem.Next
	}
	return errors.New("item not found")
}
