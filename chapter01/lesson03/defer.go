package main

import "fmt"

func main() {

	test6()

}

// de语句会延迟函数的执行直到上层函数返回。
// 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
func test5() {
	defer fmt.Println("world")
	fmt.Println("hello,")
}

// 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
func test6() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}
