package main

import (
	"fmt"
	"sort"
)

func median(i []int) float64 {
	sort.Ints(i)
	if len(i)%2 == 0 {
		return float64(i[len(i)/2]+i[len(i)/2-1]) / 2
	}
	return float64(i[len(i)/2])

}

func main() {
	// for first example my array equals 2,1,3,5,4 and result must be 3
	nums := []int{2, 1, 7, 5, 4}
	res := median(nums)
	fmt.Printf("median for %d equals %g\n", nums, res)

	// for second example my array equals 2,5,1,7 and result must be 4
	nums = []int{2, 5, 1, 7}
	res = median(nums)
	fmt.Printf("median for %d equals %g\n", nums, res)
}
