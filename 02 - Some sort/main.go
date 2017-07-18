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
	studygroup := people{"John", "Mariia", "Alex", "Petya", "Nepetya", "Pavel", "Anton"}
	s := []string{"John", "Mariia", "Alex", "Petya", "Nepetya", "Pavel", "Anton"}
	sort.Sort(studygroup)
	fmt.Println(studygroup)
	fmt.Println(s)
}
