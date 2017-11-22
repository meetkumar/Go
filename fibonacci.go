package main

import (
	"fmt"
	"strconv")

func fib(n int) int{
	var a[6] int;
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

func main(){
	fmt.Print(strconv.Itoa(fibonacci(5))+" ")
	fmt.Print(fib(5))
}
