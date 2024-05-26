package cycling

import (
	"fmt"
	"math"
	"testing"
)

func TestRideTime(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	cases := []testCase{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{}, 0},
		{[]int{1}, 1}}
	for _, tc := range cases {
		got := RideTime(&tc.test)
		if got != tc.want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestFunctionalThreshold(t *testing.T) {
	// create a new PowerMetrics type
	var metric []int
	//build sample data set.
	for i := 0; i < 2400; i++ {
		if i%3 == 0 {
			metric = append(metric, 260)
		} else if i%2 == 0 {
			metric = append(metric, 240)
		} else {
			metric = append(metric, 250)
		}
	}
	want := 237 // The above loop should build a []int that results in a 237 FT
	got := FunctionalThreshold(&metric)
	// compare against known FT result.
	if want != got {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestAverage(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	tests := []testCase{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
		{[]int{5, 2, 3, 4, 5, 6, 7, 9}, 5},
		{[]int{}, 0}}
	for _, tc := range tests {
		got := Average(&tc.test)
		if got != tc.want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestNormalized(t *testing.T) {
	type testCase struct {
		min  int
		mid  int
		max  int
		want int
	}
	tests := []testCase{
		{min: 0, mid: 0, max: 0, want: 0},
		{min: 250, mid: 250, max: 250, want: 249}, // this is weird...
		{min: 200, mid: 220, max: 250, want: 222},
		{min: 210, mid: 220, max: 230, want: 219},
		{min: 200, mid: 240, max: 300, want: 246}}
	for _, tc := range tests {
		got := Normalized(func() *[]int {
			var r []int
			for i := 0; i < 1300; i++ {
				if i%3 == 0 {
					r = append(r, tc.max)
				} else if i%2 == 0 {
					r = append(r, tc.min)
				} else {
					r = append(r, tc.mid)
				}
			}
			return &r
		}())
		want := tc.want
		if got != want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestVariabilityIndex(t *testing.T) {
	type testCase struct {
		norm int
		avg  int
		want float64
	}
	tests := []testCase{
		{norm: 250, avg: 250, want: 1.0},
		{norm: 260, avg: 240, want: 1.08333},
		{norm: 240, avg: 260, want: 0.92308},
		{norm: 0, avg: 1, want: 0.00},
		{norm: 0, avg: 0, want: math.NaN()},
		{norm: 1, avg: 0, want: math.Inf(1)}}
	for _, tc := range tests {
		got := VariabilityIndex(tc.norm, tc.avg)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}

func TestIntensityFactor(t *testing.T) {
	type testCase struct {
		norm int
		ft   int
		want float64
	}
	tests := []testCase{
		{norm: 250, ft: 250, want: 1.0},
		{norm: 270, ft: 250, want: 1.08},
		{norm: 190, ft: 250, want: 0.76},
		{norm: 250, ft: 200, want: 1.25},
		{norm: 0, ft: 1, want: 0.00000},
		{norm: 0, ft: 0, want: math.NaN()},
		{norm: 1, ft: 0, want: math.Inf(1)}}
	for _, tc := range tests {
		got := IntensityFactor(tc.norm, tc.ft)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}

func TestTrainingStressScore(t *testing.T) {
	type testCase struct {
		time int
		norm int
		ft   int
		inf  float64
		want float64
	}
	tests := []testCase{
		{time: 3600, norm: 250, ft: 250, inf: 1.0, want: 100.0},
		{time: 4800, norm: 250, ft: 250, inf: 1.0, want: 133.33333},
		{time: 3600, norm: 270, ft: 250, inf: 1.0, want: 108.0},
		{time: 3600, norm: 250, ft: 200, inf: 1.25, want: 156.25},
		{time: 1200, norm: 250, ft: 250, inf: 1.0, want: 33.33333},
		{time: 600, norm: 450, ft: 250, inf: 1.8, want: 54.0},
	}
	for _, tc := range tests {
		got := TrainingStressScore(tc.time, tc.norm, tc.ft, tc.inf)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}
