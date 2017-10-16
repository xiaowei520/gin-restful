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

//注意事项
//nil判断 尾节点开辟进位空间 刚开始赋值开辟空间
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位值, 只可能为0或1
	promotion := 0

	// 结果表的头结点
	var head *ListNode
	// 保存结果表的尾结点
	var rear *ListNode
	for nil != l1 || nil != l2 {
		sum := 0
		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			sum += l2.Val
			l2 = l2.Next
		}
		//是否进位相加
		sum += promotion
		promotion = 0
		//获取进位
		if sum > 9 {
			promotion = 1
			sum %= 10
		}
		//赋值
		node := &ListNode{
			sum,
			nil,
		}
		//头结点赋值
		if nil == head {
			head = node
			rear = node
		} else {
			rear.Next = node
			rear = node
		}

	}
	//进位判断-额外开辟空间
	if promotion > 0 {
		rear.Next = &ListNode{
			promotion,
			nil,
		}
	}

	return head
}
func main() {

	var l1 ListNode;
	var next1 *ListNode;

	//	next1.Val = 2;

	l1.Val = 1;
	l1.Next = next1;
	fmt.Println(l1);
	fmt.Println(next1);

	//l1 := ListNode{1,new()};

	//addTwoNumbers();
	fmt.Print(isPalindrome(-2147447412));
}

