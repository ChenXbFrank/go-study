package main

import "fmt"

func main(){
	result, res := max(2,5)
	fmt.Println(result, res)
}
/*func function_name( [parameter list] ) [return_types] {
	函数体
}*/
func max(num1, num2 int) (int ,string) {
	/* 声明局部变量 */
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result, "返回成功"
}
