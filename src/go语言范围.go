package main

import "fmt"

func main() {
	nums := []int{2,3,4,5}
	sum := 0
	for index,num :=range nums{
		fmt.Println(index)
		sum += num
	}
	fmt.Println(sum)
}
