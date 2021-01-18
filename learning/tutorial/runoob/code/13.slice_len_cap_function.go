/**
 * @date    2021-01-18 15:59:12
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// len=3 cap=5 slice=[0 0 0]
