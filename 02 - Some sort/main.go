/*

 */

package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var studygroup people
	s := []string{"John", "Mariia", "Alex", "Petya", "Nepetya", "Pavel", "Anton"}
	i := []int{4, 3, 1, 7, 4, 1, 8, 9, 4, 23, 4}
	studygroup = s
	fmt.Println("Original slice: \n", studygroup)
	sort.Sort(sort.StringSlice(s))
	fmt.Println("Sorted slice:\n ", s)
	sort.Sort(sort.Reverse(studygroup))
	fmt.Println("people type sorted reversely:\n ", studygroup)
	sort.Sort(sort.IntSlice(i))
	fmt.Println("Sorted int slice:\n ", i)
}
