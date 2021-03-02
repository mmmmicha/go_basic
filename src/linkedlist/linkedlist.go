package main

import "fmt"

// LinkedList
// 노드에서 노드를 포인터로 연결하는것 (node struct)
// 메모리가 빈틈없이 연결되어있는 배열과는 다름

// root pointer 만 가지고 있을 경우 node 를 추가할때 root 부터 끝까지 찾아야하는 O(N)
// root, tail pointer 를 가지고 있을 경우 node 를 추가할때 tail만 찾으면 되므로 O(1)

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)
	// root 바로 다음 노드 지우기
	root, tail = RemoveNode(root.next, root, tail)

	PrintNodes(root)
	// root 노드 지우기
	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)
	// tail 노드 지우기
	root, tail = RemoveNode(tail, root, tail)

	PrintNodes(root)
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func AddNode(tail *Node, val int) *Node {
	/*
		var tail *Node
		tail = root
		for tail.next != nil {
			tail = tail.next
		}
	*/

	node := &Node{val: val}
	tail.next = node
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	// 지우고자 하는 node 가 root 인 경우
	if node == root {
		root = root.next
		// 만약 root가 nil 인 경우 그렇다면 til 도 nil 이어야 한다.
		if root == nil {
			tail = nil
		}
		return root, tail
	}
	// root에서부터 지우고자 하는 node 앞까지 pointer 전진하기
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	// 지우고자 하는 node가 tail 인 경우
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}
