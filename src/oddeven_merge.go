package main

import (
	"sync"
)

func oddeven_merge(arr []int, lo int, n int) {
	if n > 1 {
		m := n / 2
		oddeven_merge(arr, lo, m)
		oddeven_merge(arr, lo+m, m)
		merge(arr, lo, n, 1)
	}
}
func oddeven_merge_go(arr []int, lo, n int, sem chan struct{}) {
	if n > 1 {
		m, wg := n/2, sync.WaitGroup{}
		wg.Add(2)
		select {
		case sem <- struct{}{}:
			if n >= 4096 {
				go func() {
					oddeven_merge_go(arr, lo, m, sem)
					<-sem
					wg.Done()
				}()
			} else {
				oddeven_merge(arr, lo, m)
				wg.Done()
			}
		default:
			oddeven_merge(arr, lo, m)
			wg.Done()
		}
		select {
		case sem <- struct{}{}:
			if n >= 4096 {
				go func() {
					oddeven_merge_go(arr, lo+m, m, sem)
					<-sem
					wg.Done()
				}()
			} else {
				oddeven_merge(arr, lo+m, m)
				wg.Done()
			}
		default:
			oddeven_merge(arr, lo+m, m)
			wg.Done()
		}
		wg.Wait()
		merge(arr, lo, n, 1)
	}
}
func merge(arr []int, lo, n, r int) {
	m := r * 2
	if m < n {
		merge(arr, lo, n, m)
		merge(arr, lo+r, n, m)
		for i := lo + r; i+r < lo+n; i += m {
			compare(arr, i, i+r)
		}
	} else {
		compare(arr, lo, lo+r)
	}
}
func compare(arr []int, a, b int) {
	if arr[a] > arr[b] {
		arr[a], arr[b] = arr[b], arr[a]
	}
}
