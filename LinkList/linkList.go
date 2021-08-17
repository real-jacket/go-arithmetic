package linkList

type ListNode struct {
	value int
	next  *ListNode
}

func Generate(list []int) ListNode {
	head := ListNode{value: list[0]}

	pre := &head

	for i := 1; i < len(list); i++ {
		next := ListNode{value: list[i]}
		pre.next = &next
		pre = &next
	}

	return head

}
