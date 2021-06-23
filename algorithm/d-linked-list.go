package algorithm

import (
	"fmt"
)

type dLinkedList struct {
	Data  int
	Index int
	Prev  *dLinkedList
	Next  *dLinkedList
}

func NewNode(x, index int) *dLinkedList {
	return &dLinkedList{
		Data:  x,
		Index: index,
	}
}

func (d *dLinkedList) AddNode(x, index int) {
	node := NewNode(x, index)
	d.Next = node
	node.Prev = d
}

func (d *dLinkedList) Print() {
	cur := d
	for cur != nil {
		fmt.Printf("%d:%d ", cur.Index, cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}

func (d *dLinkedList) Len() (len int) {
	cur := d
	for cur != nil {
		len++
		cur = cur.Next
	}
	return
}

func (d *dLinkedList) Get(index int) *dLinkedList {
	if index > d.Len() && index < 0 {
		return nil
	}
	cur := d
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

func (d *dLinkedList) doQuickSort(left, right *dLinkedList) {
	if left.Index <= right.Index {
		pivot := d.Get((left.Index + right.Index) / 2).Data
		i, j := left, right
		for i.Index <= j.Index {
			for i.Data < pivot {
				i = i.Next
			}
			fmt.Println(j.Data)
			for j.Data > pivot {
				j = j.Prev
			}
			if i.Index <= j.Index {
				i.Data, j.Data = j.Data, i.Data
				i = i.Next
				j = j.Prev
			}
			d.Print()
		}
		if left.Index < j.Index {
			d.doQuickSort(left, j)
		}
		if right.Index > i.Index {
			d.doQuickSort(i, right)
		}

	}
}

func (d *dLinkedList) QuickSortDLL() {
	n := d.Len() - 1

	d.doQuickSort(d, d.Get(n))
}

func GenDLL(size, min, max int) *dLinkedList {
	nums := Generate(size, min, max)
	head := NewNode(0, -1)
	cur := head
	for i := 0; i < len(nums); i++ {
		cur.AddNode(nums[i], i)
		cur = cur.Next
	}
	head = head.Next
	head.Prev = nil
	return head
}
