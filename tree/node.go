package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(Value int) {
	node.Value = Value
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}
