package cycling

import (
	"reflect"
	"testing"
)

func TestNewPowerMetrics_StaticFTP(t *testing.T) {
	ftp := 250
	pes := sampleDataPowerEachSec(2000, 200)
	pow_mets := NewPowerMetrics(ftp, pes)
	if ftp != pow_mets.FTP {
		t.Fatalf(`PowerMetrics.FTP = %d, want %d`, pow_mets.FTP, ftp)
	}
	if reflect.DeepEqual(pes, pow_mets.PowerEachSec) == false {
		t.Fatalf("PowerMetrics.PowerEachSec does not equal to value passed to power_each_second value passed to NewPowerMetrics.")
	}
	if pow_mets.Time == 0 {
		t.Fatal("Time value not initialized.")
	}
	if pow_mets.AP == 0 {
		t.Fatal("AP value not initialized.")
	}
	if pow_mets.NP == 0 {
		t.Fatal("NP value not initialized.")
	}
	if pow_mets.VI == 0 {
		t.Fatal("VI value not initialized.")
	}
	if pow_mets.IF == 0 {
		t.Fatal("IF value not initialized.")
	}
	if pow_mets.TSS == 0 {
		t.Fatal("TSS value not initialized.")
	}
}

func TestNewPowerMetrics_CalculatedFTP(t *testing.T) {
	pes := sampleDataPowerEachSec(2000, 200)
	pow_mets := NewPowerMetrics(0, pes)
	if pow_mets.FTP == 0 {
		t.Fatal("FTP value not initialized when passed as 0 to newPowerMetrics.")
	}
	if reflect.DeepEqual(pes, pow_mets.PowerEachSec) == false {
		t.Fatal("PowerMetrics.PowerEachSec does not equal to value passed to power_each_second value passed to NewPowerMetrics.")
	}
	if pow_mets.Time == 0 {
		t.Fatal("Time value not initialized.")
	}
	if pow_mets.AP == 0 {
		t.Fatal("AP value not initialized.")
	}
	if pow_mets.NP == 0 {
		t.Fatal("NP value not initialized.")
	}
	if pow_mets.VI == 0 {
		t.Fatal("VI value not initialized.")
	}
	if pow_mets.IF == 0 {
		t.Fatal("IF value not initialized.")
	}
	if pow_mets.TSS == 0 {
		t.Fatal("TSS value not initialized.")
	}
}

///////////////////////////////////////////////////////////////////////////////
// testing helper funcs

// This function just builds a large int slice to mimic DataPowerEachSec,
// since those slices are normally thousands of elements long and
// unwieldy to type statically.
func sampleDataPowerEachSec(size, min_val int) []int {
	const step_size = 9
	slice_size := size + step_size - (size % step_size)
	var dpes []int
	for i := 0; i < slice_size; i++ {
		dpes = append(dpes, min_val+(i%step_size))
	}
	return dpes
}
