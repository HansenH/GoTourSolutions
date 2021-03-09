package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	ans := 1.0
	for last := 0.0; math.Abs(ans-last) > 0.0000001; ans -= (ans*ans - x) / (2 * ans) {
		last = ans
		fmt.Println(ans)
	}
	return ans
}

func main() {
	fmt.Println(Sqrt(2))
}
