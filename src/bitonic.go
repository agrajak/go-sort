package main

func bitonic(up bool, arr []int, lo int, n int) {
	if n > 1 {
		m := n / 2
		bitonic(true, arr, lo, m)
		bitonic(false, arr, lo+m, m)
		bitonic_merge(up, arr, lo, n)
	}
}

func bitonic_merge(up bool, arr []int, lo int, n int) {
	if n > 1 {
		m := n / 2
		for i := lo; i < lo+m; i++ {
			if arr[i] > arr[i+m] == up {
				arr[i], arr[i+m] = arr[i+m], arr[i]
			}
		}
		bitonic_merge(up, arr, lo, m)
		bitonic_merge(up, arr, lo+m, m)
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
		m := n / 2
		for i := lo; i < lo+m; i++ {
			if arr[i] > arr[i+m] == up {
				arr[i], arr[i+m] = arr[i+m], arr[i]
			}
		}
		c1 := make(chan int)
		c2 := make(chan int)
		go bitonic_merge_parallel(up, arr, lo, m, c1)
		go bitonic_merge_parallel(up, arr, lo+m, m, c2)
		<-c1
		<-c2
	}
	c <- 1
}
