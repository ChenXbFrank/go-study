package main

import "fmt"

func main() {
	fmt.Println("切片初始化")
	s := []int{1,2,3,4}
	fmt.Println(s[2:])


	var numbers = make([]int,3,5)

	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

