package main

import (
	"fmt"
	"math"
)

//https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	zcd := len(nums1) + len(nums2)
	i := 0
	index1 := 0
	result := []int{}

	temp := []int{}

	lenNums1 := len(nums1)
	lenNums2 := len(nums2)

	if lenNums1 == 0 {
		result = nums2
		goto cal
	} else if lenNums2 == 0 {
		result = nums1
		goto cal
	}
	if lenNums1 < lenNums2 {
		temp = nums1
		nums1 = nums2
		nums2 = temp

	}
	for _, v := range nums1 {
	t:
		if v < nums2[i] {
			index1++
			result = append(result, v)
			if index1 == len(nums1) && i < len(nums2) {
				other := nums2[i:len(nums2)]
				for _, vvv := range other {
					result = append(result, vvv)
				}
			}
			continue
		} else {
			result = append(result, nums2[i])
			if len(nums2)-1 > i {
				i++
				goto t
			} else {
				nums2[i] = math.MaxInt64
				goto t
			}
		}
	}
cal:
	if zcd%2 == 0 {
		return (float64)(result[zcd/2-1]+result[zcd/2]) / float64(2.0)
	}

	return (float64)(result[zcd/2])
}

func main() {
	a := []int{}
	b := []int{3}
	fmt.Println(findMedianSortedArrays(a, b))
}
