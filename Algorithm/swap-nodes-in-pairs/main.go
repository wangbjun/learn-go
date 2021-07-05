package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3
func main() {
	list := newList(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)

	list.print()

	//l := reverseList(list)
	l := swapNodesInPairs(list)

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

func swapNodesInPairs(head *ListNode) *ListNode {
	list := &ListNode{Next: head}
	for prev := list; prev.Next != nil && prev.Next.Next != nil; {
		prev, prev.Next, prev.Next.Next, prev.Next.Next.Next = prev.Next, prev.Next.Next, prev.Next.Next.Next, prev.Next
	}
	return list.Next
}
