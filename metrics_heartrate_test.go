package cycling

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"
)

func TestNewHeartRateMetrics(t *testing.T) {
	var want_m HeartRateMetrics
	var got_m HeartRateMetrics
	// Unmarshal test json data into a new HeartRateMetrics type.
	td, err := os.ReadFile("testdata/metrics_heartrate_sampledata.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &want_m)
	// Save testing data metrics to compare against.
	got_m = NewHeartRateMetrics(want_m.FTHR, want_m.HeartRateEachSec)
	if want_m.Time != got_m.Time {
		t.Fatalf("Want %d, got %d", want_m.Time, got_m.Time)
	}
	if want_m.AHR != got_m.AHR {
		t.Fatalf("Want %d, got %d", want_m.AHR, got_m.AHR)
	}
	if want_m.NHR != got_m.NHR {
		t.Fatalf("Want %d, got %d", want_m.NHR, got_m.NHR)
	}
	if want_m.VI != got_m.VI {
		t.Fatalf("Want %f, got %f", want_m.VI, got_m.VI)
	}
	if want_m.INF != got_m.INF {
		t.Fatalf("Want %f, got %f", want_m.INF, got_m.INF)
	}
	if want_m.TSS != got_m.TSS {
		t.Fatalf("Want %f, got %f", want_m.TSS, got_m.TSS)
	}

}
func TestNewHeartRateMetrics_NoFTHR(t *testing.T) {
	var want_m HeartRateMetrics
	var got_m HeartRateMetrics
	// Unmarshal test json data into a new HeartRateMetrics type.
	td, err := os.ReadFile("testdata/metrics_heartrate_sampledata.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &want_m)
	// Save testing data metrics to compare against.
	got_m = NewHeartRateMetrics(0, want_m.HeartRateEachSec)
	if got_m.FTHR != 121 {
		t.Fatalf("got %d, want %d", got_m.FTHR, 121)
	}
}

func TestSessionTime_HeartRate(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	cases := []testCase{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{}, 0},
		{[]int{1}, 1}}
	var m HeartRateMetrics
	for _, tc := range cases {
		got := m.SessionTime(&tc.test)
		if got != tc.want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestFunctionalThresholdHeartRate(t *testing.T) {
	// create HeartRateMetrics type.
	var m HeartRateMetrics
	// build sample data set.
	for i := 0; i < 2400; i++ {
		if i%3 == 0 {
			m.HeartRateEachSec = append(m.HeartRateEachSec, 150)
		} else if i%2 == 0 {
			m.HeartRateEachSec = append(m.HeartRateEachSec, 130)
		} else {
			m.HeartRateEachSec = append(m.HeartRateEachSec, 140)
		}
	}
	// Calculate FTHR
	got := m.FunctionalThresholdHeartRate(&m.HeartRateEachSec)
	// The above loop should build a []int that results in a 133 FTHR
	want := 133
	// compare against know FTHR result.
	if got != want {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestFunctionalThresholdHearRate(t *testing.T) {
	// create a new HeartRateMetrics type
	var m HeartRateMetrics
	//build sample data set.
	for i := 0; i < 2400; i++ {
		if i%3 == 0 {
			m.HeartRateEachSec = append(m.HeartRateEachSec, 260)
		} else if i%2 == 0 {
			m.HeartRateEachSec = append(m.HeartRateEachSec, 240)
		} else {
			m.HeartRateEachSec = append(m.HeartRateEachSec, 250)
		}
	}
	want := 237 // The above loop should build a []int that results in a 237 fthr
	// Run the functional thresholdpower method
	got := m.FunctionalThresholdHeartRate(&m.HeartRateEachSec)
	// compare against known fthr result.
	if want != got {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestAverageHeartRate(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	tests := []testCase{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
		{[]int{5, 2, 3, 4, 5, 6, 7, 9}, 5},
		{[]int{}, 0}}
	var m HeartRateMetrics
	for _, tc := range tests {
		got := m.AverageHeartRate(&tc.test)
		if got != tc.want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestNormalizedHeartRate(t *testing.T) {
	type testCase struct {
		min  int
		mid  int
		max  int
		want int
	}
	tests := []testCase{
		{min: 0, mid: 0, max: 0, want: 0},
		{min: 120, mid: 120, max: 120, want: 119}, // this is weird...
		{min: 100, mid: 120, max: 150, want: 122},
		{min: 110, mid: 120, max: 130, want: 119},
		{min: 100, mid: 140, max: 180, want: 139}}
	var m HeartRateMetrics
	for _, tc := range tests {
		got := m.NormalizedHeartRate(func() *[]int {
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

func TestVariabilityIndex_HeartRate(t *testing.T) {
	type testCase struct {
		nhr  int
		ahr  int
		want float64
	}
	tests := []testCase{
		{nhr: 140, ahr: 140, want: 1.0},
		{nhr: 160, ahr: 140, want: 1.14286},
		{nhr: 140, ahr: 160, want: 0.87500},
		{nhr: 0, ahr: 1, want: 0.00},
		{nhr: 0, ahr: 0, want: math.NaN()},
		{nhr: 1, ahr: 0, want: math.Inf(1)}}
	var m HeartRateMetrics
	for _, tc := range tests {
		got := m.VariabilityIndex(tc.nhr, tc.ahr)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}

func TestIntensityFactor_HeartRate(t *testing.T) {
	type testCase struct {
		nhr  int
		fthr int
		want float64
	}
	tests := []testCase{
		{nhr: 140, fthr: 140, want: 1.0},
		{nhr: 170, fthr: 150, want: 1.13333},
		{nhr: 110, fthr: 150, want: 0.73333},
		{nhr: 150, fthr: 100, want: 1.5},
		{nhr: 0, fthr: 1, want: 0.00000},
		{nhr: 0, fthr: 0, want: math.NaN()},
		{nhr: 1, fthr: 0, want: math.Inf(1)}}
	var m HeartRateMetrics
	for _, tc := range tests {
		got := m.IntensityFactor(tc.nhr, tc.fthr)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}

func TestTrainingStressScore_HeartRate(t *testing.T) {
	type testCase struct {
		time int
		nhr  int
		fthr int
		inf  float64
		want float64
	}
	tests := []testCase{
		{time: 3600, nhr: 150, fthr: 150, inf: 1.0, want: 100.0},
		{time: 4800, nhr: 150, fthr: 150, inf: 1.0, want: 133.33333},
		{time: 3600, nhr: 170, fthr: 150, inf: 1.0, want: 113.33333},
		{time: 3600, nhr: 150, fthr: 100, inf: 1.25, want: 187.5},
		{time: 1200, nhr: 150, fthr: 150, inf: 1.0, want: 33.33333},
		{time: 600, nhr: 200, fthr: 150, inf: 1.8, want: 40.0},
	}
	var m HeartRateMetrics
	for _, tc := range tests {
		got := m.TrainingStressScore(tc.time, tc.nhr, tc.fthr, tc.inf)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}
