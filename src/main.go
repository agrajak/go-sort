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

var counter int

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

func nearest2Pow(n int) int {
	k := 1
	for {
		k *= 2
		if k >= n {
			return k - n
		}
	}
}
func sort(arr []int, ans []int, idx int) time.Duration {
	N := len(arr)
	input := make([]int, N, N)
	copy(arr, input)
	var t time.Time
	var du time.Duration
	m := 0
	counter = 0
	sem := make(chan struct{}, 4)
	t = time.Now()

	if idx >= 3 { // sort가 bitonic or odd-even-merge 일때
		m = nearest2Pow(N)
		input = append(input, make([]int, m)...) // 2의 지수승을 맞추기 위해서 m개의 0을 채워준다.
	}

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
	if idx >= 3 {
		input = input[m:]
	}
	du = -t.Sub(time.Now())
	fmt.Println("*", sorts[idx], "sort spend", du, verify(arr, ans))
	return du
}
func benchmark(arr []int, arg ...bool) []time.Duration {
	N := len(arr)
	result := make([]time.Duration, len(sorts))
	ans := make([]int, N, N)
	i := 0

	if len(arg) > 0 { // do not use selection sort in bench mode.
		shell(ans, N)
		i = 1 // skip selection sort
	} else {
		selection(ans, N)
	}
	for ; i < len(sorts); i++ {
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
	start, end, mul := 1, 23, 1
	mul = int(math.Pow(2, float64(start)))
	for i := 0; i < mul; i++ {
		arr = append(arr, int(rand.Int31()))
	}
	for round := start; round < end; round++ {
		fmt.Println(" *** for N =", mul, "*** by 12151616 Jeon, Suhyun")
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
