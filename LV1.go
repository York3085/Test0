package main

import "fmt"

var a int
var b int
var operator string

type FuncType func(int, int) int

func f1(a, b int) int {
	if operator == "+" {
		return a + b
	}
	if operator == "-" {
		return a - b
	}
	if operator == "*" {
		return a * b
	}
	if operator == "/" {
		return a / b
	}
	return 0
}

func main() {
	fmt.Println("请输入两个整数和一个操作符，进行四则运算")
	fmt.Println("请输入第一个整数：")
	fmt.Scanf("%d\n", &a)
	fmt.Println("请输入操作符：")
	fmt.Scanf("%s\n", &operator)
	fmt.Println("请输入第二个整数：")
	fmt.Scanf("%d\n", &b)
	var f2 FuncType = f1
	fmt.Println("计算结果是 ", f2(a, b))
}
