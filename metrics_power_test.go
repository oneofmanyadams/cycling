package cycling

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"
)

func TestNewPowerMetrics(t *testing.T) {
	var w PowerMetrics
	var got_m PowerMetrics
	var ride Ride
	ftp := 250
	// Unmarshal test json data into a new PowerMetrics type.
	td, err := os.ReadFile("testdata/sample_ride.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &ride)
	// manually call metrics functions in correct order to compare against.
	w.FTP = ftp
	w.Time = w.SessionTime(&ride.PowerEachSec)
	w.AP = w.AveragePower(&ride.PowerEachSec)
	w.NP = w.NormalizedPower(&ride.PowerEachSec)
	w.VI = w.VariabilityIndex(w.NP, w.AP)
	w.INF = w.IntensityFactor(w.NP, w.FTP)
	w.TSS = w.TrainingStressScore(w.Time, w.NP, w.FTP, w.INF)

	got_m = NewPowerMetrics(ftp, ride.PowerEachSec)
	if w.Time != got_m.Time {
		t.Fatalf("Time = %d; want %d", got_m.Time, w.Time)
	}
	if w.AP != got_m.AP {
		t.Fatalf("AP = %d, want %d", got_m.AP, w.AP)
	}
	if w.NP != got_m.NP {
		t.Fatalf("NP = %d, want %d", got_m.NP, w.NP)
	}
	if w.VI != got_m.VI {
		t.Fatalf("VI = %f, want %f", got_m.VI, w.VI)
	}
	if w.INF != got_m.INF {
		t.Fatalf("INF = %f, want %f", got_m.INF, w.INF)
	}
	if w.TSS != got_m.TSS {
		t.Fatalf("TSS = %f, want %f", got_m.TSS, w.TSS)
	}
}

func TestNewPowerMetrics_NoFTP(t *testing.T) {
	var got_m PowerMetrics
	var ride Ride
	td, err := os.ReadFile("testdata/sample_ride.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &ride)
	// Save testing data metrics to compare against.
	got_m = NewPowerMetrics(0, ride.PowerEachSec)
	if got_m.FTP != 162 {
		t.Fatalf("got: %d, want: %d", got_m.FTP, 162)
	}
}

func TestSessionTime_Power(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	cases := []testCase{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{}, 0},
		{[]int{1}, 1}}
	var m PowerMetrics
	for _, tc := range cases {
		got := m.SessionTime(&tc.test)
		if got != tc.want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestFunctionalThresholdPower(t *testing.T) {
	// create a new PowerMetrics type
	var pm PowerMetrics
	var pwrs []int
	//build sample data set.
	for i := 0; i < 2400; i++ {
		if i%3 == 0 {
			pwrs = append(pwrs, 260)
		} else if i%2 == 0 {
			pwrs = append(pwrs, 240)
		} else {
			pwrs = append(pwrs, 250)
		}
	}
	want := 237 // The above loop should build a []int that results in a 237 FTP
	// Run the functional thresholdpower method
	got := pm.FunctionalThresholdPower(&pwrs)
	// compare against known FTP result.
	if want != got {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}

func TestAveragePower(t *testing.T) {
	type testCase struct {
		test []int
		want int
	}
	tests := []testCase{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
		{[]int{5, 2, 3, 4, 5, 6, 7, 9}, 5},
		{[]int{}, 0}}
	var m PowerMetrics
	for _, tc := range tests {
		got := m.AveragePower(&tc.test)
		if got != tc.want {
			t.Fatalf("got: %d, want: %d", got, tc.want)
		}
	}
}

func TestNormalizedPower(t *testing.T) {
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
	var m PowerMetrics
	for _, tc := range tests {
		got := m.NormalizedPower(func() *[]int {
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

func TestVariabilityIndex_Power(t *testing.T) {
	type testCase struct {
		np   int
		ap   int
		want float64
	}
	tests := []testCase{
		{np: 250, ap: 250, want: 1.0},
		{np: 260, ap: 240, want: 1.08333},
		{np: 240, ap: 260, want: 0.92308},
		{np: 0, ap: 1, want: 0.00},
		{np: 0, ap: 0, want: math.NaN()},
		{np: 1, ap: 0, want: math.Inf(1)}}
	var m PowerMetrics
	for _, tc := range tests {
		got := m.VariabilityIndex(tc.np, tc.ap)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}

func TestIntensityFactor_Power(t *testing.T) {
	type testCase struct {
		np   int
		ftp  int
		want float64
	}
	tests := []testCase{
		{np: 250, ftp: 250, want: 1.0},
		{np: 270, ftp: 250, want: 1.08},
		{np: 190, ftp: 250, want: 0.76},
		{np: 250, ftp: 200, want: 1.25},
		{np: 0, ftp: 1, want: 0.00000},
		{np: 0, ftp: 0, want: math.NaN()},
		{np: 1, ftp: 0, want: math.Inf(1)}}
	var m PowerMetrics
	for _, tc := range tests {
		got := m.IntensityFactor(tc.np, tc.ftp)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}

func TestTrainingStressScore_Power(t *testing.T) {
	type testCase struct {
		time int
		np   int
		ftp  int
		inf  float64
		want float64
	}
	tests := []testCase{
		{time: 3600, np: 250, ftp: 250, inf: 1.0, want: 100.0},
		{time: 4800, np: 250, ftp: 250, inf: 1.0, want: 133.33333},
		{time: 3600, np: 270, ftp: 250, inf: 1.0, want: 108.0},
		{time: 3600, np: 250, ftp: 200, inf: 1.25, want: 156.25},
		{time: 1200, np: 250, ftp: 250, inf: 1.0, want: 33.33333},
		{time: 600, np: 450, ftp: 250, inf: 1.8, want: 54.0},
	}
	var m PowerMetrics
	for _, tc := range tests {
		got := m.TrainingStressScore(tc.time, tc.np, tc.ftp, tc.inf)
		if fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tc.want) {
			t.Fatalf("got: %.5f, want: %.5f", got, tc.want)
		}
	}
}
