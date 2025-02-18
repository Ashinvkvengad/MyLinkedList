package main

import (
	SL "MyLinkedList/singlyLinkedList"
)

func TEST0() {
	//list := new(SL.List)
	list := SL.CreateSinglyLinkedList()
	list.PrintMyList()
	list.AddNodeToBegin()
	// list.AddNodeToBegin()
	// list.AddNodeToBegin()
	// list.AddNodeToBegin()
	// list.AddNodeToBegin()
	// list.AddNodeToBegin()
	list.AddNodeToMiddle(3)
	list.AddNodeToEnd()
	list.PrintMyList()
	list.DeleteNodeFromList(2, SL.NODE_INDEX)
	list.DeleteNodeFromList(2, SL.NODE_ID)
	list.PrintMyList()
}

func main() {
	TEST0()
}
