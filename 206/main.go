package main


//节点
type ListNode struct {
	Val int
	Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
	var (
		//cur初始值为nil，为返回的头节点
		//next为中转节点
		cur,next *ListNode
	)

	for head!=nil{
		//先保存下一个
		next = head.Next
		//把上一个赋值给当前的下一个
		head.Next = cur
		//把当前的赋值给cur
		cur = head
		//把下一个赋值给head
		head = next
	}
	return cur
}

func main()  {
	
}


