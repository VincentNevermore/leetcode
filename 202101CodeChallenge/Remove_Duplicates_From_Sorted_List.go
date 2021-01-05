// 2021 Jan Week 1 Challenge
// Author: Shengwei Zhang

/**
 * Definition for singly-linked list.
 */

class ListNode {
    val: number
    next: ListNode | null

    constructor(val?: number, next?: ListNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}

function deleteDuplicates(head: ListNode | null): ListNode | null {
    if (head == null) {
        return null
    }
    
};