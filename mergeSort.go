package main

import "fmt"

func mergeSort(arr []int) []int{
	if len(arr) <=1{
	 	return arr
	}
	mid := len(arr)/2 
	mergeSort(arr[:mid])
	mergeSort(arr[mid:])
	return merge(arr,mid)
}

func merge(arr []int, mid int) []int{
	L := make([]int, len(arr[:mid]))
	R := make([]int, len(arr[mid:]))
	copy(L,arr[:mid])
	copy(R,arr[mid:])
	m := len(L)
	n := len(R)
	var i,j,k int
	for i<m && j<n{
		if L[i] < R[j]{
			arr[k] = L[i]
			i++
		}else{
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i<m{
		arr[k] = L[i]
		k++
		i++
	}
	for j<n{
		arr[k] = R[j]
		k++
		j++
	}
	return arr 
}

func main(){
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	mergeSort(s)
	fmt.Println(s)
}
