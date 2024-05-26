package cycling

import (
	"encoding/json"
	"os"
	"testing"
)

func TestNewHeartRateMetrics(t *testing.T) {
	var w HeartRateMetrics
	var got_m HeartRateMetrics
	var ride Ride
	fthr := 121
	// Unmarshal test json data into a new HeartRateMetrics type.
	td, err := os.ReadFile("testdata/sample_ride.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &ride)
	// Manually call metrics functions in correct order to compate against.
	w.FTHR = fthr
	w.Time = RideTime(&ride.HeartRateEachSec)
	w.AHR = Average(&ride.HeartRateEachSec)
	w.NHR = Normalized(&ride.HeartRateEachSec)
	w.VI = VariabilityIndex(w.NHR, w.AHR)
	w.INF = IntensityFactor(w.NHR, w.FTHR)
	w.TSS = TrainingStressScore(w.Time, w.NHR, w.FTHR, w.INF)

	got_m = NewHeartRateMetrics(fthr, ride.HeartRateEachSec)
	if w.Time != got_m.Time {
		t.Fatalf("Want %d, got %d", w.Time, got_m.Time)
	}
	if w.AHR != got_m.AHR {
		t.Fatalf("Want %d, got %d", w.AHR, got_m.AHR)
	}
	if w.NHR != got_m.NHR {
		t.Fatalf("Want %d, got %d", w.NHR, got_m.NHR)
	}
	if w.VI != got_m.VI {
		t.Fatalf("Want %f, got %f", w.VI, got_m.VI)
	}
	if w.INF != got_m.INF {
		t.Fatalf("Want %f, got %f", w.INF, got_m.INF)
	}
	if w.TSS != got_m.TSS {
		t.Fatalf("Want %f, got %f", w.TSS, got_m.TSS)
	}
}

func TestNewHeartRateMetrics_NoFTHR(t *testing.T) {
	var got_m HeartRateMetrics
	var ride Ride
	// Unmarshal test json data into a new HeartRateMetrics type.
	td, err := os.ReadFile("testdata/sample_ride.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &ride)
	// Save testing data metrics to compare against.
	got_m = NewHeartRateMetrics(0, ride.HeartRateEachSec)
	if got_m.FTHR != 121 {
		t.Fatalf("got %d, want %d", got_m.FTHR, 121)
	}
}
