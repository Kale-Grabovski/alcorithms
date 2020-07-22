package main

import "fmt"

func merge(a []int, b []int) (ret []int) {
	ai, bi, al, bl := 0, 0, len(a), len(b)

	for ai < al || bi < bl {
		if ai < al && bi < bl {
			if b[bi] > a[ai] {
				ret = append(ret, a[ai])
				ai++
				continue
			}
		} else if ai < al {
			ret = append(ret, a[ai])
			ai++
			continue
		}

		ret = append(ret, b[bi])
		bi++
	}

	return ret
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	return merge(mergeSort(a[0:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func main() {
	fmt.Println(mergeSort([]int{5, 1, 0, 10, 2}))
	fmt.Println(mergeSort([]int{15, 5, 1, 0, 10, 2}))
}
