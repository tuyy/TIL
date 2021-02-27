package book

import (
	"container/list"
	"fmt"
	"testing"
)

func TestNo16(t *testing.T) {
	l1, l2 := list.New(), list.New()
	l1.PushBack(2)
	l1.PushBack(4)
	l1.PushBack(3)

	l2.PushBack(5)
	l2.PushBack(6)
	l2.PushBack(4)

	result := addTwoNumbers(l1.Front(), l2.Front())
	//l3 := list.New()
	//n1 := l1.Front()
	//n2 := l2.Front()

	//add := 0
	//for ; n1 != nil && n2 != nil; n1, n2 = n1.Next(), n2.Next() {
	//	sum := n1.Value.(int) + n2.Value.(int) + add
	//	if sum < 10 {
	//		l3.PushBack(sum)
	//		add = 0
	//	} else {
	//		l3.PushBack(sum - 10)
	//		add = 1
	//	}
	//}

	//for n1 != nil {
	//	l3.PushBack(n1.Value.(int) + add)
	//	add = 0
	//	n1 = n1.Next()
	//}
	//for n2 != nil {
	//	l3.PushBack(n2.Value.(int) + add)
	//	add = 0
	//	n2 = n2.Next()
	//}

	//if add != 0 {
	//	l3.PushBack(1)
	//}

	for n := result.Front(); n != nil; n = n.Next() {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}

func addTwoNumbers(n1, n2 *list.Element) *list.List {
	result := list.New()

	carry := 0

	for n1 != nil || n2 != nil || carry == 1 {
		sum := carry
		if n1 != nil {
			sum += n1.Value.(int)
			n1 = n1.Next()
		}
		if n2 != nil {
			sum += n2.Value.(int)
			n2 = n2.Next()
		}

		if sum < 10 {
			carry = 0
		} else {
			sum -= 10
			carry = 1
		}

		result.PushBack(sum)
	}
	return result
}
