/**
 * @date    2021-01-18 10:48:15
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

/* 交换函数这样写更加简洁，也是 go 语言的特性，可以用下，c++ 和 c# 是不能这么干的 */

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
