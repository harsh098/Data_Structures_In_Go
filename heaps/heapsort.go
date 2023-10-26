package main

func HeapSort(array *[]int){
	h := new(MaxHeap)
	h.array = (*array)[:]
	h.MaxHeapify(0)
	for i:=len(*array)-1 ;len(h.array) > 0; i--{
		(*array)[i] = h.Extract()
	}
}