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
	var num1, num2 int
	var carry int
	res := &ListNode{Val: 0}
	look := res
	for (l1 != nil || l2 != nil) || carry == 1 {
		if l1 != nil {
			num1 = l1.Val
		} else {
			num1 = 0
		}
		if l2 != nil {
			num2 = l2.Val
		} else {
			num2 = 0
		}
		look.Val = (num1 + num2 + carry) % 10
		if num1+num2+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 != nil || l2 != nil || carry == 1 {
			next := &ListNode{Val: 0}
			look.Next = next
			look = look.Next
		}
	}
	return res
}
