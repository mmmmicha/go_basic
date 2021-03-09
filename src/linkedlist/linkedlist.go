package main

import "fmt"

// LinkedList
// 노드에서 노드를 포인터로 연결하는것 (node struct)
// 메모리가 빈틈없이 연결되어있는 배열과는 다름

// root pointer 만 가지고 있을 경우 node 를 추가할때 root 부터 끝까지 찾아야하는 O(N)
// root, tail pointer 를 가지고 있을 경우 node 를 추가할때 tail만 찾으면 되므로 O(1)

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	// 단순 복제이기 때문에 아래 구문은 현재 l.tail 의 pointer 를 순수하게 갖게됨
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = node.next
		l.root.prev = nil
		node.next = nil
		return
	}

	// root에서부터 지우고자 하는 node 앞까지 pointer 전진하기
	// prev := l.root
	// for prev.next != node {
	// 	prev = prev.next
	// }

	prev := node.prev

	// 지우고자 하는 node가 tail 인 경우
	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
	node.next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
	// root 바로 다음 노드 지우기
	list.RemoveNode(list.root.next)

	list.PrintNodes()
	// root 노드 지우기
	list.RemoveNode(list.root)

	list.PrintNodes()
	// tail 노드 지우기
	list.RemoveNode(list.tail)

	list.PrintNodes()

	fmt.Println("*******************")
	list.PrintReverse()
	fmt.Println("*******************")

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

	// sclice vs linkedlist
	// 이 경우 length:5, capacity:5
	// append 를 하더라도 capacity:5 유지, length만 4로 바뀜
	// 고로 새로운 메모리를 할당하진 않음
	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a)
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
