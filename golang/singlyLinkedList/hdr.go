package singlyLinkedList

const (
	NODE_INDEX = 10
	NODE_ID    = 11
)

type Node struct {
	Data   interface{}
	NodeId uint
	Next   *Node
}

type List struct {
	Head  *Node
	count uint
}
