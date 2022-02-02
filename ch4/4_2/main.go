package main

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func MinimalTree(arr []int, start int, end int) *TreeNode {
	if end < start {
		return nil
	}

	mid := (start + end) / 2

	//create node with mid
	n := TreeNode{
		data: arr[mid],
	}
	n.left = MinimalTree(arr, start, mid-1)
	n.right = MinimalTree(arr, mid+1, end)
	return &n
}

func AMinimalTree(arr []int) *TreeNode {
	return MinimalTree(arr, arr[0], arr[len(arr)-1])
}

func main() {

}
