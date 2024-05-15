package cycling

// ToDo
// -implement table-driven tests where applicable.
import (
	"encoding/json"
	"os"
	"testing"
)

func TestNewPowerMetrics(t *testing.T) {
	var want_pm PowerMetrics
	var got_pm PowerMetrics
	// Unmarshal test json data into a new PowerMetrics type.
	td, err := os.ReadFile("testdata/metrics_power_sampledata.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &want_pm)
	// Save testing data metrics to compare against.
	got_pm = NewPowerMetrics(want_pm.FTP, want_pm.PowerEachSec)
	t.Run("Calculate Time", func(t *testing.T) {
		if want_pm.Time != got_pm.Time {
			t.Fatalf("Want %d, got %d", want_pm.Time, got_pm.Time)
		}
	})
	t.Run("Calculate AP", func(t *testing.T) {
		if want_pm.AP != got_pm.AP {
			t.Fatalf("Want %d, got %d", want_pm.AP, got_pm.AP)
		}
	})
	t.Run("Calculate NP", func(t *testing.T) {
		if want_pm.NP != got_pm.NP {
			t.Fatalf("Want %d, got %d", want_pm.NP, got_pm.NP)
		}
	})
	t.Run("Calculate VI", func(t *testing.T) {
		if want_pm.VI != got_pm.VI {
			t.Fatalf("Want %f, got %f", want_pm.VI, got_pm.VI)
		}
	})
	t.Run("Calculate INF", func(t *testing.T) {
		if want_pm.INF != got_pm.INF {
			t.Fatalf("Want %f, got %f", want_pm.INF, got_pm.INF)
		}
	})
	t.Run("Calculate TSS", func(t *testing.T) {
		if want_pm.TSS != got_pm.TSS {
			t.Fatalf("Want %f, got %f", want_pm.TSS, got_pm.TSS)
		}
	})

}
