package main

func median_of_three(arr []int, left int, right int) {
	mid := (left + right) / 2
	// sort between (left, mid, right)
	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[mid] > arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if right-left+1 > 3 {
		pivot := arr[mid]
		arr[mid], arr[right-1] = arr[right-1], arr[mid]
		i, j := left, right-1
		for {
			for {
				i++
				if i >= right || arr[i] >= pivot {
					break
				}
			}
			for {
				j--
				if left >= j || arr[j] <= pivot {
					break
				}
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i], arr[right-1] = arr[right-1], arr[i]
		median_of_three(arr, left, i-1)
		median_of_three(arr, i+1, right)
	}
}
