package main

import "fmt"

func main() {
	fmt.Println("定义数组")
	var n [10]int
	var i,j int
	for i=0;i< 10;i++{
		n[i] = i + 100
	}

	for j =0;j<len(n);j++{
		fmt.Println(n[j])
	}
}
