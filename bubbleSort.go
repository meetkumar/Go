package main

import "fmt"

func sort(arr [10]int){
	var sorted =true
	for sorted{
		sorted = false
		for i:=0;i<len(arr)-1;i++{
			if(arr[i+1] < arr[i]){
				arr[i],arr[i+1] = arr[i+1],arr[i]
				sorted = true
			}
		}
	}	
	fmt.Println(arr)
}

func main(){
	var toBeSorted [10]int = [10]int{1,3,2,4,8,6,7,2,3,0}
	fmt.Println("arr : ",toBeSorted)
	sort(toBeSorted)
	//fmt.Println("Sorted array : ",toBeSorted)
}
