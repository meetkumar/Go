package main

import "fmt"

func main(){
a:=[]int{1,2,3}

for i, arr := range a{
	if arr%2 == 0{
		fmt.Println("a: ",i)
	}		
}


kvs := map[string]string{"a":"apple", "b":"banana"}

for i, k:= range kvs{
	fmt.Println("key: ",i ," value: ",k)
}


}
