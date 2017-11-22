package main

import "fmt"
import "math"

const s string ="constant"

func main(){
	a:="asshole"
	fmt.Println(s)
	const n=5000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Print(a)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
