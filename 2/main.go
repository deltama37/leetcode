package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	c := &ListNode{
		Val: 0,
	}
	b := &ListNode{
		Val:  4,
		Next: c,
	}
	a := &ListNode{
		Val:  2,
		Next: b,
	}
	f := &ListNode{
		Val: 8,
	}
	e := &ListNode{
		Val:  1,
		Next: f,
	}
	d := &ListNode{
		Val:  5,
		Next: e,
	}
	// addTwoNumbers(a, d)
	fmt.Println(a.Val, a.Next.Val, a.Next.Next.Val)
	fmt.Println(d.Val, d.Next.Val, d.Next.Next.Val)
	res := addTwoNumbers(e, c)
	fmt.Println(res.Val, res.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var node *ListNode
	var n int
	for l1 != nil || l2 != nil || n > 0 {
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}
		node = &ListNode{n % 10, node}
		n /= 10
	}
	var newNode *ListNode
	for node != nil {
		newNode = &ListNode{node.Val, newNode}
		node = node.Next
	}
	return newNode
}
