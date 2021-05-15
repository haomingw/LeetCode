/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public var val: Int
 *     public var next: ListNode?
 *     public init() { self.val = 0; self.next = nil; }
 *     public init(_ val: Int) { self.val = val; self.next = nil; }
 *     public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
 * }
 */
class Solution {
  func addTwoNumbers(_ l1: ListNode?, _ l2: ListNode?) -> ListNode? {
    var dummy = ListNode(0)
    var prev = dummy
    var l1 = l1
    var l2 = l2
    var carry = 0
    while l1 != nil || l2 != nil || carry > 0 {
      let value = (l1 == nil ? 0 : l1!.val) + (l2 == nil ? 0 : l2!.val) + carry
      let cur = ListNode(value % 10)
      carry = value / 10
      prev.next = cur
      prev = cur
      l1 = l1?.next
      l2 = l2?.next
    }
    return dummy.next
  }
}
