package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// selection
// median-of-three
// shell
// bitonic
// odd-even merge
func verify(a []int, b []int) bool {
	N := len(a)
	for i := 0; i < N; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func compute(original []int) {
	N := len(original)
	var input []int = make([]int, N)
	var ans []int = make([]int, N)
	var t time.Time
	var du time.Duration
	//var c chan int

	//	fmt.Printf("before sort ... :")
	//fmt.Print(original, "\n")
	copy(input, original)
	t = time.Now()
	selection(input, N)
	du = -t.Sub(time.Now())
	copy(ans, input)
	//fmt.Print("after selection sort ... :", input, "\n")
	fmt.Print("selection spend ", du, " => ", verify(ans, input), "\n")

	copy(input, original)
	t = time.Now()
	median_of_three(input, 0, N-1)
	du = -t.Sub(time.Now())
	//	fmt.Print("after median-of-three quick sort ... :", input, "\n")
	fmt.Print("median-of-three spend ", du, " => ", verify(ans, input), "\n")

	copy(input, original)
	t = time.Now()
	shell(input, N)
	du = -t.Sub(time.Now())
	//	fmt.Print("after shell ... :", input, "\n")
	fmt.Print("shell spend ", du, " => ", verify(ans, input), "\n")

	copy(input, original)
	t = time.Now()
	bitonic(true, input, 0, N)
	du = -t.Sub(time.Now())
	//	fmt.Print("after bitonic ... :", input, "\n")
	fmt.Print("bitonic spend ", du, " => ", verify(ans, input), "\n")

	//c = make(chan int)
	sem := make(chan struct{}, 2)
	copy(input, original)
	t = time.Now()
	bitonic_parallel(true, input, 0, N, sem)

	//<-c
	du = -t.Sub(time.Now())
	//	fmt.Print("after (parallel) bitonic ... :", input, "\n")
	fmt.Print("bitonic(goroutine) spend ", du, " => ", verify(ans, input), "\n")
}

func main() {
	var array []int

	start, end, mul := 18, 19, 1
	mul = int(math.Pow(2, float64(start)))
	for i := 0; i < end-start; i++ {
		for j := 0; j < mul; j++ {
			array = append(array, int(rand.Int31()))
		}
		compute(array)
		mul *= 2
	}

}
