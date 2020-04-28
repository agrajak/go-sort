package main

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func selection(arr []int, N int) {
	for i := 0; i < N; i++ {
		minIdx := i
		for j := i; j < N; j++ {
			// find min in j for i
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}
