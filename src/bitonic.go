package main

func bitonic(up bool, arr []int, lo int, n int) {
	if n > 1 {
		bitonic(true, arr, lo, n/2)
		bitonic(true, arr, lo+n/2, n/2)
		bitonic_merge(up, arr, lo, n)
	}
}

func bitonic_merge(up bool, arr []int, lo int, n int) {
	if n > 1 {
		dist := len(arr) / 2
		for i := 0; i < dist; i++ {
			if arr[i] > arr[i+dist] == up {
				arr[i], arr[i+dist] = arr[i+dist], arr[i]
			}
		}
		bitonic_merge(up, arr, lo, n/2)
		bitonic_merge(up, arr, lo+n/2, n/2)
	}
}

func bitonic_parallel(up bool, arr []int, lo int, n int, c chan int) {
	if n > 1 {
		c1 := make(chan int)
		c2 := make(chan int)
		go bitonic_parallel(true, arr, lo, n/2, c1)
		go bitonic_parallel(true, arr, lo+n/2, n/2, c2)
		<-c1
		<-c2
		bitonic_merge_parallel(up, arr, lo, n, c)
	}
	c <- 1
	/*
		first := bitonic(true, arr[:len(arr)/2], pFlag)
		second := bitonic(false, arr[len(arr)/2:], pFlag)
		return bitonic_merge(up, append(first, second...), pFlag)
	*/
}

func bitonic_merge_parallel(up bool, arr []int, lo int, n int, c chan int) {
	if n > 1 {
		// compare and swap
		dist := len(arr) / 2
		for i := 0; i < dist; i++ {
			if arr[i] > arr[i+dist] == up {
				arr[i], arr[i+dist] = arr[i+dist], arr[i]
			}
		}
		c1 := make(chan int)
		c2 := make(chan int)
		go bitonic_merge_parallel(up, arr, lo, n/2, c1)
		go bitonic_merge_parallel(up, arr, lo+n/2, n/2, c2)
		<-c1
		<-c2
	}
	c <- 1
}
