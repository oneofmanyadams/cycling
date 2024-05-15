package cycling

// ToDo
// -implement table-driven tests where applicable.
import (
	"encoding/json"
	"os"
	"testing"
)

func TestNewHeartRateMetrics(t *testing.T) {
	var want_m HeartRateMetrics
	var got_m HeartRateMetrics
	// Unmarshal test json data into a new PowerMetrics type.
	td, err := os.ReadFile("testdata/metrics_heartrate_sampledata.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &want_m)
	// Save testing data metrics to compare against.
	got_m = NewHeartRateMetrics(want_m.FTHR, want_m.HeartRateEachSec)
	t.Run("Calculate Time", func(t *testing.T) {
		if want_m.Time != got_m.Time {
			t.Fatalf("Want %d, got %d", want_m.Time, got_m.Time)
		}
	})
	t.Run("Calculate AHR", func(t *testing.T) {
		if want_m.AHR != got_m.AHR {
			t.Fatalf("Want %d, got %d", want_m.AHR, got_m.AHR)
		}
	})
	t.Run("Calculate NHR", func(t *testing.T) {
		if want_m.NHR != got_m.NHR {
			t.Fatalf("Want %d, got %d", want_m.NHR, got_m.NHR)
		}
	})
	t.Run("Calculate VI", func(t *testing.T) {
		if want_m.VI != got_m.VI {
			t.Fatalf("Want %f, got %f", want_m.VI, got_m.VI)
		}
	})
	t.Run("Calculate INF", func(t *testing.T) {
		if want_m.INF != got_m.INF {
			t.Fatalf("Want %f, got %f", want_m.INF, got_m.INF)
		}
	})
	t.Run("Calculate TSS", func(t *testing.T) {
		if want_m.TSS != got_m.TSS {
			t.Fatalf("Want %f, got %f", want_m.TSS, got_m.TSS)
		}
	})

}
