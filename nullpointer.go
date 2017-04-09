package main

import "fmt"

func main() {
	var ptr *int
	
	if ptr == nil {
		fmt.Printf("ptr是空指针！")
		fmt.Printf("ptr的值为：%x\n", ptr)
	}else{
		fmt.Printf("ptr不是空指针！")
		fmt.Printf("ptr的值为：%x\n", ptr)
	}
	
}