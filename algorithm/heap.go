package algorithm

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Pop() (x interface{}) {
	n := len(*h)
	x = (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	(*h) = append((*h), x.(int))
}
