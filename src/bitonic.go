package main

import (
	"sync"
)

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

func bitonic_go(up bool, arr []int, lo int, n int, sem chan struct{}) {
	if n > 1 {
		wg := sync.WaitGroup{}
		wg.Add(2)
		select {
		case sem <- struct{}{}:
			go func() {
				bitonic_go(true, arr, lo, n/2, sem)
				<-sem
				wg.Done()
			}()
		default:
			bitonic(true, arr, lo, n/2)
			wg.Done()
		}
		select {
		case sem <- struct{}{}:
			go func() {
				bitonic_go(false, arr, lo+n/2, n/2, sem)
				<-sem
				wg.Done()
			}()
		default:
			bitonic(false, arr, lo+n/2, n/2)
			wg.Done()
		}
		wg.Wait()

		bitonic_merge(up, arr, lo, n)
	}
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
