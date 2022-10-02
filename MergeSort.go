package main

import "fmt"

func main() {

	arr := []int{12, 11, 13, 5, 6, 7}
	l := arr[:(len(arr))/2]
	r := arr[(len(arr))/2:]
	fmt.Println(mergeSort(l, r))
}

func mergeSort(l, r []int) []int {

	arr2 := make([]int, len(l)+len(r))
	i, j, k := 0, 0, 0

	for len(l) > i && len(r) > j {
		if l[i] < r[j] {
			arr2[k] = l[i]
			i++
		} else {
			arr2[k] = r[j]
			j++
		}
		k++
	}

	// if any element was left

	for len(l) > i {
		arr2[k] = l[i]
		i++
		k++
	}

	for len(r) > j {
		arr2[k] = r[j]
		j++
		k++
	}

	return arr2
}
