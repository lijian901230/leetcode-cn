package main

import "fmt"

/**
 * 两两交换链表中的节点
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	     Val int
	     Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	curr := &ListNode{Next:head}
	var tmp = curr
	for tmp.Next != nil && tmp.Next.Next!=nil {
		//获取第一个next 和第二个next的值
		a,b := tmp.Next,tmp.Next.Next
		a.Next = b.Next
		b.Next = a
		tmp.Next = b
		tmp = a

	}
	return curr.Next
}

func swapPairs2(head *ListNode)*ListNode{
	var (
		cur,tmp *ListNode
	)

	cur = &ListNode{Next:head}
	tmp = cur

	for tmp.Next!=nil && tmp.Next.Next!=nil  {
		//获取当前第一个和第二个
		//a为第一个  b为第二个
		a,b := tmp.Next,tmp.Next.Next
		//a.Next实际为第二个 b.Next为第三个
		a.Next = b.Next
		//
		b.Next = a
		tmp.Next = b
		tmp = a
	}
	return cur.Next
}

func main()  {
	a := &ListNode{Val:4}
	b := &ListNode{Val:7,Next:a}
	c := &ListNode{Val:2,Next:b}
	d := &ListNode{Val:1,Next:c}
	e := &ListNode{Val:5,Next:d}

	//for d !=nil{
	//	fmt.Println(d.Val)
	//	d = d.Next
	//}
	F := swapPairs(e)
	fmt.Print(F)

}
