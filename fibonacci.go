package main

import (
	"fmt"
	)

func fib(n int) int{
	a := make([]int,n+1);
	a[0] = 0
	a[1] = 1
	for i:=2;i<=n;i++{
		a[i] = a[i-1] + a[i-2]
	}
	return a[n]
} 

func fibonacci(n int) int{
	if(n <= 1){ 
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2);
	
}

func fibmap(n int) int{
	if n==0{
		return 0
	}else if n<=2{
		return 1
	}
	a := make(map[int]int)
	if _,ok := a[n];ok{
		return a[n]
	}
	result:= fibmap(n-1)+fibmap(n-2)
	a[n]=result
	return result
}

func main(){
	fmt.Println("Resursive fibonacci :",fibonacci(5))
	fmt.Println("Iteratice fibonacci :",fib(5))
	fmt.Print("HashMap fibonacci :",fibmap(5))
}
