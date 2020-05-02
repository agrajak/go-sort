package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func verify(a []int, b []int) bool {
	N := len(a)
	for i := 0; i < N; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var sorts []string = []string{"selection", "median of three", "shell", "bitonic", "bitonic(go)", "oddeven merge", "oddeven merge(go)"}

func sort(arr []int, ans []int, idx int) time.Duration {
	N := len(arr)
	input := make([]int, N, N)
	copy(arr, input)
	var t time.Time
	var du time.Duration
	sem := make(chan struct{}, 4)
	t = time.Now()
	switch idx {
	case 0:
		selection(input, N)
	case 1:
		median_of_three(input, 0, N-1)
	case 2:
		shell(input, N)
	case 3:
		bitonic(true, input, 0, N)
	case 4:
		bitonic_go(true, input, 0, N, sem)
	case 5:
		oddeven_merge(input, 0, N)
	case 6:
		oddeven_merge_go(input, 0, N, sem)
	}
	du = -t.Sub(time.Now())
	fmt.Println("*", sorts[idx], "sort spend", du, verify(arr, ans))
	return du
}
func benchmark(arr []int) []time.Duration {
	N := len(arr)
	result := make([]time.Duration, len(sorts))
	ans := make([]int, N, N)
	selection(ans, N)
	for i := 0; i < len(sorts); i++ {
		result[i] = sort(arr, ans, i)
	}
	return result
}

func main() {
	var arr []int

	file, err := os.Create("./result.csv")
	if err != nil {
		panic(nil)
	}
	wr := csv.NewWriter(bufio.NewWriter(file))
	start, end, mul := 1, 22, 1
	mul = int(math.Pow(2, float64(start)))
	for i := 0; i < mul; i++ {
		arr = append(arr, int(rand.Int31()))
	}
	for round := start; round < end; round++ {
		fmt.Println(" *** for N =", mul, "***")
		paper := make([]string, len(sorts)+1)
		paper[0] = strconv.Itoa(round)
		for i, r := range benchmark(arr) {
			paper[i+1] = fmt.Sprintf("%f", float64(r/time.Microsecond))
		}
		wr.Write(paper)
		wr.Flush()
		for i := 0; i < mul; i++ {
			arr = append(arr, int(rand.Int31()))
		}
		mul *= 2
	}

}
