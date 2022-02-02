package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head Node
}

func main() {
	someList := List{
		head: Node{
			data: 8,
			next: &Node{
				data: 5,
				next: &Node{
					data: 3,
					next: &Node{
						data: 5,
						next: &Node{
							data: 4,
							next: nil,
						},
					},
				},
			},
		},
	}

	someList.RemoveDuplicateNodes()
}

func (L *List) RemoveDuplicateNodes() error {
	tracker := make(map[int]int)

	node := L.head

	if node.next == nil {
		return nil
	}

	for node.next != nil {
		fmt.Printf("%d ->", node.data)

		if tracker[node.next.data] < 1 {
			tracker[node.next.data] += 1
			node = *node.next
		} else {
			if node.next.next == nil {
				break
			} else {
				node = *node.next.next
			}

		}
	}

	fmt.Print(node.data)

	return nil

}

func (L *List) RemoveDuplicateNodesNoBuffer() error {

	node := L.head

	for node.next != nil {

		if node.data != node.next.data {
			node = node.
		}

	}

}
