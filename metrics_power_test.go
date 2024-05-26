package cycling

import (
	"encoding/json"
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
	w.Time = RideTime(&ride.PowerEachSec)
	w.AP = Average(&ride.PowerEachSec)
	w.NP = Normalized(&ride.PowerEachSec)
	w.VI = VariabilityIndex(w.NP, w.AP)
	w.INF = IntensityFactor(w.NP, w.FTP)
	w.TSS = TrainingStressScore(w.Time, w.NP, w.FTP, w.INF)

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
