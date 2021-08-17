package linkList

import "testing"

func LogList(head *ListNode, t *testing.T) {
	if head != nil {
		t.Log("node: ", head)
		LogList(head.next, t)
	}

}

func TestGnerate(t *testing.T) {
	list := []int{2, 9, 4}

	ListNode1 := Generate(list)

	LogList(&ListNode1, t)

}
