package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//l := &ListNode{Val: 2}
	l1 := &ListNode{Val: 8}
	l2 := &ListNode{Val: 1,Next:l1}

	//q := &ListNode{Val: 5}
	//q1 := &ListNode{Val: 6, Next: q}
	q2 := &ListNode{Val: 0}
	//fmt.Println(q2)
	hh := addTwoNumbers(l2, q2)
	//fmt.Println(hh)
	for hh != nil {
		fmt.Println(hh)
		hh = hh.Next
	}
}

// 2->4->3
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	falg := l
	falg2 :=0
	for {
		if l1!=nil {
			falg.Val += l1.Val
			l1 = l1.Next
		}

		if l2!=nil {
			falg.Val += l2.Val
			l2 = l2.Next
		}
		falg.Val+=falg2
		if falg.Val>=10 {
			falg.Val -= 10
			falg2 = 1
		}else {
			falg2 = 0
		}

		if l1==nil && l2==nil {
			break
		}
		falg.Next = &ListNode{}
		falg = falg.Next
	}
	if falg2 >0 {
		falg.Next = &ListNode{Val:falg2}
	}
	return l
}
