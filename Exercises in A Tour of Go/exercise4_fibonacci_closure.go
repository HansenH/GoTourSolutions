package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		ans := a
		a, b = b, b+a
		return ans
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
