package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func sortedSquares(A []int) []int {
	var result []int

	for _, v := range A {
		result = append(result, v*v)
	}
	sort.Sort(IntSlice(result))

	return result
}
func main() {
	a:=[]int{-1,-2,44,123,0}
	a=sortedSquares(a)
	fmt.Println(a)
}