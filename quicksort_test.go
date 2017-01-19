package quicksort

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type sortable []int

func (s sortable) Less(i, j int) bool { return s[i] < s[j] }
func (s sortable) Len() int           { return len(s) }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TestSimple4(t *testing.T) {
	input := []int{1, 3, 2, 4}
	correct := []int{1, 2, 3, 4}
	Sort(sortable(input)) // mutates input slice
	for i := 0; i < len(input); i++ {
		if input[i] != correct[i] {
			t.Fatal("sort incorrect... received:", input, "expected:", correct)
		}
	}
}

func TestSimple10(t *testing.T) {
	input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	correct := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Sort(sortable(input)) // mutates input slice
	for i := 0; i < len(input); i++ {
		if input[i] != correct[i] {
			t.Fatal("sort incorrect... received:", input, "expected:", correct)
		}
	}
}

func readInput(name string) []int {
	input, _ := ioutil.ReadFile(name)
	s := strings.Split(string(input), "\n")
	output := make([]int, len(s))
	for i, item := range s {
		intItem, _ := strconv.Atoi(item)
		output[i] = intItem
	}
	return output
}

func TestLargeSet(t *testing.T) {
	input := readInput("test_input.txt")
	correct := readInput("test_input.txt")
	sort.Sort(sortable(correct))
	Sort(sortable(input)) // mutates input slice
	for i := 0; i < len(input); i++ {
		if input[i] != correct[i] {
			t.Fatal("sort incorrect... received:", input, "expected:", correct)
		}
	}
}
