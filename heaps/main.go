package main

import (
	"fmt"
)

type MaxHeap struct{
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1) //Start from last element	
}

func (h* MaxHeap) maxHeapifyUp(index int){
	for h.array[h.parent(index)] < h.array[index]{
		h.array[h.parent(index)], h.array[index] = h.array[index], h.array[h.parent(index)]
		index = h.parent(index)
	}
}

func (h* MaxHeap) MaxHeapifyUsingTraditionalBuildHeap(){
	for i:=1 ; i<len(h.array) ; i++{
		h.maxHeapifyUp(i)
	}
}

func (h* MaxHeap) parent(index int) int{
	if index==0{
		return 0
	}

	if index%2==0 {
		return index/2 -1
	}

	return (index-1)/2
}

func (h* MaxHeap) Extract() int{
	if len(h.array) == 0 {
		panic("Empty Heap")
	}
	extracted := h.array[0]
	l := len(h.array)-1
	h.array[l], h.array[0] = h.array[0], h.array[l]
	h.array = h.array[:l]
	h.MaxHeapify(0)
	return extracted
} 

func (h *MaxHeap) MaxHeapify(index int){
	l,r := h.left(index), h.right(index)
	largest := index
	if l<len(h.array) && h.array[l] > h.array[index] {
		largest = l
	}
	if r<len(h.array) && h.array[r] > largest {
		largest = r
	}
	if largest != index {
		h.array[largest], h.array[index] = h.array[index], h.array[largest]
		h.MaxHeapify(largest)
	}
}

func (h* MaxHeap) left(index int) int{
	return 2*index + 1
}

func (h* MaxHeap) right(index int) int{
	return 2*index + 2
}


func main()  {
	m := &MaxHeap{[]int{10,20,30,15}}
	m.MaxHeapifyUsingTraditionalBuildHeap()
	fmt.Println(m)
	for len(m.array) > 0 {
		x := m.Extract()
		fmt.Println(m, ",", x)
	}
	fmt.Println(m)
	
}