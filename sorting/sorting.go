package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func RunSorting() {
	alphas := []string{"c", "a", "e"}
	sort.Strings(alphas)
	fmt.Printf("alphas: %v\n", alphas)

	isSorted := sort.StringsAreSorted(alphas)
	fmt.Printf("isSorted: %v\n", isSorted)

	numbers := rand.Perm(10)
	sort.Ints(numbers)
	key := 2
	index := sort.SearchInts(numbers, key)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("i: %v\n", index)
}
