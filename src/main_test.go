package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestMain(t *testing.T) {
	var arr []int

	start, end, mul := 18, 19, 1
	mul = int(math.Pow(2, float64(start)))
	for i := 0; i < mul; i++ {
		arr = append(arr, int(rand.Int31()))
	}
	for round := start; round < end; round++ {
		fmt.Println(" *** for N =", mul, "*** by 12151616 Jeon, Suhyun")
		benchmark(arr, true)
		for i := 0; i < mul; i++ {
			arr = append(arr, int(rand.Int31()))
		}
		mul *= 2
	}

}
