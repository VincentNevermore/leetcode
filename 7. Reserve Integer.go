package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	result := 0
	if x < 0 {
		sign = -1
		x = x * sign
	}
	for x > 0 {
		remainder := x % 10
		result *= 10
		result += remainder
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}

	return result * sign

}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
}
