package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
func main() {
	list := newList(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)

	list.print()

	l := reverseList(list)

	l.print()
}

func newList(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (list *ListNode) append(val int) {
	for list.Next != nil {
		list = list.Next
	}
	list.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (list *ListNode) print() {
	for {
		fmt.Printf("%d\n", list.Val)
		if list.Next == nil {
			break
		}
		list = list.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var cur = head
	var pre *ListNode
	for {
		cur.Next, pre, cur = pre, cur, cur.Next
		if cur == nil {
			break
		}
	}
	return pre
}
