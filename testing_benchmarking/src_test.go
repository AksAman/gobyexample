package testingbenchmarking

import (
	"fmt"
	"testing"
)

func compare(t *testing.T, got, want any) {
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestIntMinBasic(t *testing.T) {
	got := IntMin(2, -2)
	want := -2
	compare(t, got, want)
}

func TestIntMinTableDriven(t *testing.T) {
	var testCases = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, testCase := range testCases {
		testName := fmt.Sprintf("Testing IntMin(%d, %d)=%d", testCase.a, testCase.b, testCase.want)
		t.Run(
			testName,
			func(t *testing.T) {
				got := IntMin(testCase.a, testCase.b)
				compare(t, got, testCase.want)
			},
		)
	}
}

func BenchmarkIntMin(b *testing.B) {
	fmt.Println("b.N", b.N)
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
