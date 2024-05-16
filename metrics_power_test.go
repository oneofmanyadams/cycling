package cycling

// ToDo
// -implement table-driven tests where applicable.
// ---Could use multiple seeds for building the loop that for FTP test
// -Individual test for each method now that they accept args and return vals.
import (
	"encoding/json"
	"os"
	"testing"
)

func TestNewPowerMetrics(t *testing.T) {
	var want_m PowerMetrics
	var got_m PowerMetrics
	// Unmarshal test json data into a new PowerMetrics type.
	td, err := os.ReadFile("testdata/metrics_power_sampledata.json")
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(td, &want_m)
	// Save testing data metrics to compare against.
	got_m = NewPowerMetrics(want_m.FTP, want_m.PowerEachSec)
	if want_m.Time != got_m.Time {
		t.Fatalf("Time = %d; want %d", got_m.Time, want_m.Time)
	}
	if want_m.AP != got_m.AP {
		t.Fatalf("AP = %d, want %d", got_m.AP, want_m.AP)
	}
	if want_m.NP != got_m.NP {
		t.Fatalf("NP = %d, want %d", got_m.NP, want_m.NP)
	}
	if want_m.VI != got_m.VI {
		t.Fatalf("VI = %f, want %f", got_m.VI, want_m.VI)
	}
	if want_m.INF != got_m.INF {
		t.Fatalf("INF = %f, want %f", got_m.INF, want_m.INF)
	}
	if want_m.TSS != got_m.TSS {
		t.Fatalf("TSS = %f, want %f", got_m.TSS, want_m.TSS)
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
	//build sample data set.
	for i := 0; i < 2400; i++ {
		if i%3 == 0 {
			pm.PowerEachSec = append(pm.PowerEachSec, 260)
		} else if i%2 == 0 {
			pm.PowerEachSec = append(pm.PowerEachSec, 240)
		} else {
			pm.PowerEachSec = append(pm.PowerEachSec, 250)
		}
	}
	want := 237 // The above loop should build a []int that results in a 237 FTP
	// Run the functional thresholdpower method
	got := pm.FunctionalThresholdPower(&pm.PowerEachSec)
	// compare against known FTP result.
	if want != got {
		t.Fatalf("got: %d, want: %d", got, want)
	}
}
