package sorting

import (
	"fmt"
	"sort"
)

/*
*
A custom type has to be defined with implementing the

	sort.Interface - Len, Less, and Swap

*
*/
type families [][]string
type famSortedByChildren families

func (fam famSortedByChildren) Len() int {
	return len(fam)
}

func (fam famSortedByChildren) Swap(i, j int) {
	fam[i], fam[j] = fam[j], fam[i]
}

func (fam famSortedByChildren) Less(i, j int) bool {
	return len(fam[i]) < len(fam[j])
}

func RunSortingFunctions() {
	fams := families{
		{
			"a", "b", "c",
		},
		{
			"q", "w", "e", "r", "t", "y",
		},
		{
			"w", "a", "s", "d",
		},
	}

	// let's say we want to sort these families by number of
	// children
	sort.Sort(famSortedByChildren(fams))
	fmt.Printf("family: %v\n", fams)
}
