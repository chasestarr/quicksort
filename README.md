## [Quicksort](https://en.wikipedia.org/wiki/Quicksort)

[![GoDoc](https://godoc.org/github.com/chasestarr/quicksort?status.svg)](https://godoc.org/github.com/chasestarr/quicksort)

```go
import (
  "fmt"

  "github.com/chasestarr/quicksort"
)

// implements same interface as Go's 'sort' package
type sortable []int

func (s sortable) Less(i, j int) bool { return s[i] < s[j] }
func (s sortable) Len() int           { return len(s) }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
  input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
  quicksort.Sort(sortable(input))
  fmt.Println(input) // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
}
```