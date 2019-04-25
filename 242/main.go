package main



 type ListNode struct {
	     Val int
	     Next *ListNode
	 }

func reverseList(head *ListNode) *ListNode {
	var (
		curr,flag *ListNode
	)
	for head != nil {
		//先保存下个ListNode
		flag = head.Next
		//设置当前遍历节点的Next节点为curr
		head.Next = curr
		//把当前遍历的节点赋值给curr
		curr = head
		head = flag
	}
	return curr
}

func main()  {
	//reverseList()
}

