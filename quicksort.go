package quicksort

// Interface for collections, based on go std Sort
// https://golang.org/src/sort/sort.go?s=471:749#L2
type Interface interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

// Sort sorts an array of ints according to the quicksort algorithm
func Sort(a Interface) {
	quicksort(a, 0, a.Len())
}

func quicksort(a Interface, left, right int) {
	// base case
	if left >= right {
		return
	}

	// recursive cases
	p := partition(a, left, right)
	quicksort(a, left, p-1)
	quicksort(a, p, right)
}

func partition(a Interface, left, right int) int {
	// currently returns first index of subset,
	// will refactor to more thoughtful approach
	i := left + 1
	for j := left + 1; j < right; j++ {
		if a.Less(j, left) {
			a.Swap(i, j)
			i++
		}
	}
	a.Swap(i-1, left)
	return i
}
