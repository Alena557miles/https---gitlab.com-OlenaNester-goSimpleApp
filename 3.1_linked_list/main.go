package main

import "fmt"

type node struct {
	value string
	next  *node
}
type linkedList struct {
	head   *node
	lenght int
}

func (l *linkedList) append(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

func (l linkedList) printListNodes() {
	toPrint := l.head
	for l.lenght != 0 {
		fmt.Print(" ", toPrint.value)
		toPrint = toPrint.next
		l.lenght--
	}

}

func main() {
	list := linkedList{}
	node1 := &node{value: "first"}
	node2 := &node{value: "second"}
	node3 := &node{value: "third"}
	node4 := &node{value: "fourth"}
	node5 := &node{value: "fifth"}
	list.append(node1)
	list.append(node2)
	list.append(node3)
	list.append(node4)
	list.append(node5)
	fmt.Println(list)
	list.printListNodes()
}
