package cycling

import (
	"testing"
)

func TestMinAvgInts(t *testing.T) {
	type testCase struct {
		ints []int
		min  int
		want float64
	}
	cases := []testCase{
		{ints: []int{3, 3, 3}, min: 3, want: 3.0},
		{ints: []int{3, 3, 3}, min: 6, want: 1.5},
		{ints: []int{3, 3, 3}, min: 2, want: 3.0},
		{ints: []int{1, 2, 3}, min: 3, want: 2.0},
		{ints: []int{}, min: 0, want: 0.0},
	}
	for k, tc := range cases {
		got := minAvgInts(tc.ints, tc.min)
		if got != tc.want {
			t.Fatalf("Test %d got: %f, want: %f", k, got, tc.want)
		}
	}
}

func TestSumInts(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	cases := []testCase{
		{test: []int{1, 2, 3}, want: 6},
		{test: []int{0, 0, 0}, want: 0},
		{test: []int{5}, want: 5},
	}
	for k, tc := range cases {
		got := sumInts(tc.test)
		if got != tc.want {
			t.Fatalf("Test %d got: %d, want: %d", k, got, tc.want)
		}
	}
}
