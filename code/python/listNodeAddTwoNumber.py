import unittest

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def addTwoNumbers(l1: ListNode, l2: ListNode) -> ListNode:
    dummy = ListNode(0) # 虚拟头节点
    current = dummy
    carry = 0       # 进位

    while l1 or l2 or carry:
        val1 = l1.val if l1 else 0
        val2 = l2.val if l2 else 0

        total = val1 + val2 + carry
        carry = total // 10
        current_digit = total % 10

        current.next = ListNode(current_digit)
        current = current.next

        if l1: l1 = l1.next
        if l2: l2 = l2.next

    return dummy.next


def create_linked_list(lst):
    dummy = ListNode(0)
    current = dummy
    for val in lst:
        current.next = ListNode(val)
        current = current.next

    return dummy.next

def linked_list_to_list(head):
    result = []
    current = head
    while current:
        result.append(current.val)
        current = current.next

    return result

class TestAddTwoNumbers(unittest.TestCase):
    def test_example1(self):
        l1 = create_linked_list([2, 4, 3])
        l2 = create_linked_list([5, 6, 4])
        result = addTwoNumbers(l1, l2)
        self.assertEqual(linked_list_to_list(result), [7, 0, 8])

    def test_example2(self):
        l1 = create_linked_list([9, 9, 9])
        l2 = create_linked_list([9, 9])
        result = addTwoNumbers(l1, l2)
        self.assertEqual(linked_list_to_list(result), [8, 9, 0, 1])

    def test_zero(self):
        l1 = create_linked_list([0])
        l2 = create_linked_list([0])
        result = addTwoNumbers(l1, l2)
        self.assertEqual(linked_list_to_list(result), [0])

    def test_carry_over(self):
        l1 = create_linked_list([5])
        l2 = create_linked_list([5])
        result = addTwoNumbers(l1, l2)
        self.assertEqual(linked_list_to_list(result), [0, 1])


if __name__ == '__main__':
    unittest.main()
