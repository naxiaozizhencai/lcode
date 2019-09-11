package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


var node1 ListNode
func main(){

	node5 := ListNode{10, nil}
	node4 := ListNode{9, &node5}
	node3 := ListNode{1, &node4}
	node2 := ListNode{5, &node3}
	node1 = ListNode{4, &node2}
	deleteNode(&node3)

}

func deleteNode(node *ListNode) {
	for {

		fmt.Println(node1.Val)
		return
	}
}