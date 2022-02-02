package main

import "reflect"

func main() {

}

type State int64

const (
	Unvisited State = iota
	Visited
	Visiting
)

type Graph struct {
	start Node
}

type Node struct {
	data  int
	state State
	next  *Node //used for queue tracker
	edges []Node
	prev  *Node
}

type LinkedList struct {
	head Node
	tail Node
}

func Search(g Graph, start Node, end Node) bool {
	if reflect.DeepEqual(start, end) {
		return true
	}

	//Operates as a Queue
	q := LinkedList{}

	for _, n := range g.GetNodes() {
		n.state = Unvisited
	}

	start.state = Visited
	q.AddToList(start)

	var u Node

	for !q.IsEmpty() {
		u = q.RemoveFirst() // i.e., dequeue
		if &u.data != nil {
			for _, n := range u.edges {
				if n.state == Unvisited {
					if reflect.DeepEqual(n, end) {
						return true
					} else {
						n.state = Visiting
						q.AddToList(n)
					}
				}
			}
		}
		u.state = Visited
	}
	return false
}
