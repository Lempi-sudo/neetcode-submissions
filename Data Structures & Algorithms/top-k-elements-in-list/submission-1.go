
type NumFreq struct {
	num  int
	freq int
}

type MinHeap []NumFreq

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(NumFreq))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topK(nums map[int]int, k int) []NumFreq {
	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range nums {
		if h.Len() < k {
			heap.Push(h, NumFreq{num: num, freq: freq})
		} else if freq > (*h)[0].freq {
			heap.Pop(h)
			heap.Push(h, NumFreq{num: num, freq: freq})
		}
	}

	return *h
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num] += 1
	}
	f := topK(freq, k)
	res := make([]int, k)
	for i, v := range f {
		res[i] = v.num
	}
	return res
}