package generics

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	keysList := make([]K, 0, len(m))
	for k := range m {
		keysList = append(keysList, k)
	}
	return keysList
}

type node[T any] struct {
	next *node[T]
	val  T
}

type DLList[T any] struct {
	head *node[T]
	tail *node[T]
}

func (lst *DLList[T]) Push(v T) {
	newNode := &node[T]{val: v}
	if lst.tail == nil {
		lst.head = newNode
		lst.tail = lst.head
	} else {
		lst.tail.next = newNode
		lst.tail = lst.tail.next
	}
}

func (lst *DLList[T]) GetAllVals() []T {
	var vals []T
	for e := lst.head; e != nil; e = e.next {
		vals = append(vals, e.val)
	}
	return vals
}

func Run() {

	m := map[int]string{1: "one", 2: "two"}
	fmt.Printf("MapKeys[int](m): %v\n", MapKeys(m))

	lst := DLList[string]{}

	lst.Push("Hello")
	lst.Push("World!")

	fmt.Printf("lst.GetAllVals(): %v\n", lst.GetAllVals())
}
