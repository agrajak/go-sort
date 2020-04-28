package main

func gap_insertion(arr []int, first int, last int, K int) {
	var i, j int
	for i = first + K; i <= last; i += K {
		key := arr[i]
		for j = i - K; j >= first && arr[j] > key; j -= K {
			arr[j+K] = arr[j]
		}
		arr[j+K] = key
	}
}
func shell(arr []int, N int) {
	K := N / 2
	for {
		for i := 0; i < K; i++ {
			gap_insertion(arr, i, N-1, K)
		}
		if K == 1 {
			break
		}
		K = K / 2
		if K%2 == 0 {
			K++
		}
	}
}
