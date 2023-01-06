/*
19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]

https://leetcode.com/problems/remove-nth-node-from-end-of-list/
*/

package main

type ListNode struct {
     Val int
     Next *ListNode
 }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, delay := dummy, dummy
	for fast.Next != nil {
			fast = fast.Next
			if (n < 1) {
					delay = delay.Next
			}
			n -= 1
	}
	delay.Next = delay.Next.Next
	return dummy.Next
}
	