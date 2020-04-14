package main

import (
	"fmt"
	"sort"
)

//Function that returns an average value of array
func average(arr [6]int) float64 {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	return float64(sum) / 6
}

//Function that returns the longest word from the slice of strings(the first if there are more than one).
func max(s []string) string {
	maxValue := ""
	for _, v := range s {
		if len(maxValue) < len(v) {
			maxValue = v
		}
	}
	return maxValue
}

//Function that returns the copy of the original slice in reverse order.
func reverse(s []int64) []int64 {
	s1 := make([]int64, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		s1[len(s1)-i-1] = s[i]
	}
	return s1
}

//Function that print map values sorted in order of increasing keys.
func printSorted(m map[int]string) {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	msg := ""
	for _, v := range keys {
		msg = msg + " " + m[v]
	}

	fmt.Printf("Input: %v\n", m)
	fmt.Println("Map values sorted in order of increasing keys:" + msg)
}

func main() {

	//Task 1. Need to return the average value of the array.
	nums := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("in array %v average value equals %g\n", nums, average(nums))

	//Task 2. Need to return the longest word from the slice of strings.
	//Case 1: one of the three strings is larger than the others
	s := []string{"one", "two", "three"}
	fmt.Printf("in slice %v longest word is %s\n", s, max(s))

	//Case 2: two strings are equal. The result should be the first of these stings.
	s = []string{"one", "two"}
	fmt.Printf("in slice %v longest word is %s\n", s, max(s))

	//Task 3. Need to return the copy of the original slice in reverse order.
	s1 := []int64{1, 2, 5, 15}
	fmt.Printf("Original slice: %v\nReversed slice: %v\n", s1, reverse(s1))

	//Task 4. Need to print map values sorted in order of increasing keys.
	mp1 := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(mp1)

	mp2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(mp2)
}
