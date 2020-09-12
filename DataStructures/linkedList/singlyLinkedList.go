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

// AddAtHead adds a node at the beginning(head) of the Linked List
func (ll *linkedList) AddAtHead(data interface{}) bool {
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

// AddAtTail adds an element at the end(tail) of the list.
func (ll *linkedList) AddAtTail(data interface{}) bool {
	node := &node{
		Data: data,
	}
	if ll.Head == nil {
		ll.Head = node
		ll.Length++
		return true
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	ll.Length++
	return true
}

// AddAtIndex adds a node at given index in the Linked List. Returns an error if the index is invalid.
//Note: The list is zero indexed. Therefore the provided index must fall in the range 0 <= index <= SizeOfList.
func (ll *linkedList) AddAtIndex(index int, data interface{}) error {
	node := &node{
		Data: data,
	}
	if ll.Head == nil {
		// ll.Head = node
		return errors.New("invalid index. List is empty")
	}
	if index > ll.Size() || index < 0 {
		return errors.New("invalid index. index must be a number between 0 and list.Size()")
	}
	if index == 0 {
		ll.AddAtHead(data)
		return nil
	}
	if index == ll.Size() {
		ll.AddAtTail(data)
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

// Find searches for the specified item in the list. If found, returns the index of the item, else returns an error. In case of duplicates, returns the index of the first occurence.
func (ll *linkedList) Find(data interface{}) (int, error) {

	if ll.Head == nil {
		return 0, errors.New("item not found. List is empty")
	}
	index := 0
	currentItem := ll.Head
	if data == currentItem.Data {
		return 0, nil
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

// Get returns the value at the specified index. Returns an error if the list is empty or the index is invalid.
func (ll *linkedList) Get(index int) (interface{}, error) {
	if ll.Head == nil {
		return nil, errors.New("list is empty")
	}
	if index < 0 || index >= ll.Length {
		return nil, errors.New("invalid index")
	}
	if index == 0 {
		return ll.Head.Data, nil
	}
	current := ll.Head
	for index > 0 {
		current = current.Next
		index--
	}
	return current.Data, nil
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

// RemoveFromHead deletes the first element(Head) from the list. Returns an error if the list is empty.
func (ll *linkedList) RemoveFromHead() (interface{}, error) {
	if ll.Head == nil {
		return nil, errors.New("list is empty")
	}
	currentItem := ll.Head
	ll.Head = ll.Head.Next
	currentItem.Next = nil
	ll.Length--
	return currentItem.Data, nil

}

// RemoveFromIndex deletes the item present at specified index in the list. Return an error if the index is invalid.
// Note: List is zero indexed. Provided index must be in the range 0 <= index < SizeOfList.
func (ll *linkedList) RemoveFromIndex(index int) (interface{}, error) {
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
	if currentItem.Next == nil {
		previousItem.Next = nil
	} else {
		previousItem.Next = currentItem.Next
	}
	//currentItem.Next = nil
	ll.Length--
	return currentItem.Data, nil
}

// RemoveItem deletes first occurence of the given item from the list. Returns an error if the list is empty or the item was not found.
func (ll *linkedList) RemoveItem(data interface{}) error {
	if ll.Head == nil {
		return errors.New("list is empty")
	}
	if data == ll.Head.Data {
		ll.Head = ll.Head.Next
		ll.Length--
		return nil
	}
	currentItem := ll.Head
	previousItem := ll.Head
	for currentItem.Next != nil {
		previousItem = currentItem
		currentItem = currentItem.Next
		if currentItem.Data == data {
			if currentItem.Next == nil {
				previousItem.Next = nil
			} else {
				previousItem.Next = currentItem.Next
			}
			ll.Length--
			return nil
		}
	}
	return errors.New("item not found")
}
