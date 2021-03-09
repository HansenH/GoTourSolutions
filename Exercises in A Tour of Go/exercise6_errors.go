package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(*e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, &err
	}
	ans := 1.0
	for last := 0.0; math.Abs(ans-last) > 0.0000001; ans -= (ans*ans - x) / (2 * ans) {
		last = ans
	}
	return ans, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
