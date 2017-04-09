package main

import "fmt"


func main() {
	var a int = 100
	var b int = 200
	var result int
	
	fmt.Printf("交换前 a 的值为 : %d\n", a )
	fmt.Printf("交换前 b 的值为 : %d\n", b )
	
	result = swap2(a, b)
	fmt.Printf("交换后结果 result 的值为 : %d\n", result);
	
	fmt.Printf("交换后 a 的值为 : %d\n", a )
	fmt.Printf("交换后 b 的值为 : %d\n", b )
}

func swap2(x, y int) int {
	var temp int
	
	temp = x 
	x = y
	y = temp
	
	return temp
}