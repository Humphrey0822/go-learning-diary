package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int
	
	fmt.Printf("�� 1 �� - a��������Ϊ = %T\n", a);
	fmt.Printf("�� 2 �� - b��������Ϊ = %T\n", b);
	fmt.Printf("�� 3 �� - c��������Ϊ = %T\n", c);
	
	ptr = &a
	fmt.Printf("a ��ֵΪ %d\n", a);
	fmt.Printf("*ptr Ϊ %d\n", *ptr);
}