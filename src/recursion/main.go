package main

import "fmt"

func main() {
	fmt.Println(fib(11))
}

// fib 斐波那契数列 获取第n个斐波那契数
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 优化1 保存之前计算过的值

// 优化2 去除递归调用 自底向上计算

// 优化3 滚动数组

// 优化4 直接两个变量

// 汉诺塔

// 尾递归相关知识点：