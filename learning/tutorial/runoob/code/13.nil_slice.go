/**
 * @date    2021-01-18 16:01:06
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

package main

import "fmt"

func main() {
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Println("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// len=0 cap=0 slice=[]
// 切片是空的
