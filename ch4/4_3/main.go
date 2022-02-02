package main

func main() {

	input := BinaryTree{}
	ListOfDepths(input)

}

type LinkedList struct {
	head TreeNode
}

func (l *LinkedList) insert(n TreeNode) {
	n.next = &l.head
	l.head = n
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	next  *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func ListOfDepths(b BinaryTree) {
	listOfDepths := make(map[int]LinkedList)
	var l LinkedList
	d := 0
	AlgListOfDepths(*b.root, l, d, listOfDepths)
}

func AlgListOfDepths(root TreeNode, l LinkedList, d int, listOfDepths map[int]LinkedList) map[int]LinkedList {
	l.insert(root)
	if root.left == nil && root.right == nil {
		listOfDepths[d] = l
		return listOfDepths
	}
	leftroot := root.left
	AlgListOfDepths(*leftroot, l, d, listOfDepths)

	rightroot := root.right
	AlgListOfDepths(*rightroot, l, d, listOfDepths)

	d++

	return nil

}
