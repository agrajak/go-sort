package main

import "sync"

func oddeven_merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := oddeven_merge(arr[:mid])
	right := oddeven_merge(arr[mid:])

	return merge(left, right)
}
func oddeven_merge_go(arr []int, sem chan struct{}) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	wg := sync.WaitGroup{}
	wg.Add(2)
	var left, right []int
	select {
	case sem <- struct{}{}:
		go func() {
			left = oddeven_merge_go(arr[:mid], sem)
			<-sem
			wg.Done()
		}()
	default:
		left = oddeven_merge(arr[:mid])
		wg.Done()
	}
	select {
	case sem <- struct{}{}:
		go func() {
			right = oddeven_merge_go(arr[mid:], sem)
			<-sem
			wg.Done()
		}()
	default:
		right = oddeven_merge(arr[mid:])
		wg.Done()
	}
	wg.Wait()

	return merge(left, right)
}
func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	arr := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			arr[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			arr[k] = left[i]
			i++
		} else if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	return arr
}
