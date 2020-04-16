package main

import (
	"fmt"
	"math/rand"
)

func median(i []int) float64 {
	i = quicksort(i)
	if len(i)%2 == 0 {
		return float64(i[len(i)/2]+i[len(i)/2-1]) / 2
	}
	return float64(i[len(i)/2])

}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	// for first example my array equals 2,1,3,5,4 and result must be 3
	nums := []int{2, 1, 3, 5, 4}
	res := median(nums)
	fmt.Printf("median for %d equals %g\n", nums, res)

	// for second example my array equals 2,5,1,7 and result must be 3.5
	nums = []int{2, 5, 1, 7}
	res = median(nums)
	fmt.Printf("median for %d equals %g\n", nums, res)
}
