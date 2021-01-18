/**
 * @date    2021-01-18 10:44:42
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

/**
 * 三级指针
 */

package main

import "fmt"

func main() {
	var a int = 1
	var ptr1 *int = &a
	var ptr2 **int = &ptr1
	var ptr3 **(*int) = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2
	// 依次类推
	fmt.Println("a:", a)
	fmt.Println("ptr1", ptr1)
	fmt.Println("ptr2", ptr2)
	fmt.Println("ptr3", ptr3)
	fmt.Println("*ptr1", *ptr1)
	fmt.Println("**ptr2", **ptr2)
	fmt.Println("**(*ptr3)", **(*ptr3)) // 也可以写作：***ptr3
}

// a: 1
// ptr1 0xc000014090
// ptr2 0xc00000e028
// ptr3 0xc00000e030
// *ptr1 1
// **ptr2 1
// **(*ptr3) 1