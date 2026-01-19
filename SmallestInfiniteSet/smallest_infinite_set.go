package SmallestInfiniteSet

import "container/heap"

//2336. Smallest Number in Infinite Set

//You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].
//
//Implement the SmallestInfiniteSet class:
//
//SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
//int popSmallest() Removes and returns the smallest integer contained in the infinite set.
//void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	minVal int
	heap   *IntHeap
	added  map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		minVal: 1,
		heap:   h,
		added:  make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.heap.Len() > 0 {
		smallest := heap.Pop(this.heap).(int)
		delete(this.added, smallest)
		return smallest
	}

	// Если куча пуста, берем текущий порог и сдвигаем его
	res := this.minVal
	this.minVal++
	return res
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.minVal && !this.added[num] {
		heap.Push(this.heap, num)
		this.added[num] = true
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
