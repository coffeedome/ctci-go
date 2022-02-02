package main

func main() {

}

type Node struct {
	data   *Node
	left   *Node
	right  *Node
	parent *Node
}

func Successor(n Node) Node {

	if n.left != nil {
		Successor(*n.left)
		Successor(*n.data)
		Successor(*n.right)

	}

	return *n.data

}
