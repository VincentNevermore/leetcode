// 2021 Jan Week 1 Challenge
/**
 * Definition for singly-linked list.
*/
// 递归处理

 type ListNode struct {
      Val int
      Next *ListNode
 }

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    if head.Val == head.Next.Val {
        for head.Next != nil && head.Val ==head.Next.Val {
            head = head.Next
        }
        return deleteDuplicates(head.Next)
    }

    head.Next = deleteDuplicates(head.Next)
    return head
}