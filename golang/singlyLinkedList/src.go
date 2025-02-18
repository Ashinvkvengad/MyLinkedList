package singlyLinkedList

import (
	"fmt"
	"os"
)

var nodeIdG uint

func (list *List) listAssert(result int, msg string) {
	if 0 == result {
		fmt.Fprintf(os.Stderr, "Assersion Failed: %s", msg)
		os.Exit(1)
	}
}

func newNode() *Node {
	node := new(Node)
	node.NodeId = nodeIdG
	nodeIdG++
	return node
}

func CreateSinglyLinkedList() *List {
	list := new(List)
	node := newNode()
	list.Head = node
	list.count++
	return list
}
func (list *List) AddNodeToBegin() {
	newNode := newNode()
	newNode.Next = list.Head
	list.Head = newNode
	list.count++
}
func (list *List) AddNodeToMiddle(nodeIdx int) {
	node := list.Head
	prev := node
	idx := 0
	for nil != node {
		if nodeIdx == idx {
			break
		}
		prev = node
		node = node.Next
		idx++
	}
	newNode := newNode()
	newNode.Next = prev.Next
	prev.Next = newNode
	list.count++
}
func (list *List) AddNodeToEnd() {
	node := list.Head
	for nil != node.Next {
		node = node.Next
	}
	newNode := newNode()
	node.Next = newNode
	list.count++
}
func (list *List) DeleteNodeFromList(value uint, oper int) error {
	node := list.Head
	prev := node
	if NODE_INDEX == oper {
		if value > list.count-1 {
			return fmt.Errorf("Index should be within the range of 0 to %d", list.count-1)
		}
		idx := uint(0)
		for nil != node {
			if value == idx {
				break
			}
			prev = node
			node = node.Next
			idx++
		}
	} else if NODE_ID == oper {
		for nil != node {
			if value == node.NodeId {
				break
			}
			prev = node
			node = node.Next
		}
	}
	if nil == node {
		return fmt.Errorf("Could not find the Node from the list!")
	}
	if list.Head == prev {
		list.Head = list.Head.Next
	} else {
		prev.Next = node.Next
	}
	list.count--
	if 0 == list.count {
		list.DeleteList()
		return fmt.Errorf("There is no Nodes available in this list. The list is deleted permenently!")
	}
	return nil
}
func (list *List) DeleteList() {
	list = nil
	nodeIdG = 0
}
func (list *List) FindDataFromList() {
	//TBD
}
func (list *List) FindNodeFromList(value uint, oper int) (*Node, error) {
	node := list.Head
	if NODE_INDEX == oper {
		if value > list.count-1 {
			return nil, fmt.Errorf("Index should be within the range of 0 to %d", list.count-1)
		}
		idx := uint(0)
		for nil != node {
			if value == idx {
				return node, nil
			}
			node = node.Next
			idx++
		}
	} else if NODE_ID == oper {
		for nil != node {
			if value == node.NodeId {
				return node, nil
			}
			node = node.Next
		}
	}
	return nil, fmt.Errorf("Could not find the Node from the List!")
}
func (list *List) PrintMyList() {
	node := list.Head
	if nil == node {
		fmt.Printf("Empty List!\n")
		return
	}

	fmt.Printf("Printing the List with %d Nodes..\n", list.count)
	for nil != node {
		fmt.Printf("NodeId: %d\n", node.NodeId)
		node = node.Next
	}
}
