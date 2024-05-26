package cycling

// PowerMetrics aggregates and calculates various Power Metrics.
// The minimum input neededis PowerEachSec, all other fields can be
// populated based on that data.
// The function NewPowerMetrics() is the prefered way to generate this type.
type PowerMetrics struct {
	Time int     // in seconds
	FTP  int     // Functional Threshold Power
	AP   int     // Average Power
	NP   int     // Normalized Power
	VI   float64 // Variability Index
	INF  float64 // Intensity Factor
	TSS  float64 // Training Stress Score
}

// NewPowerMetrics is the prefered method of creating a PowerMetrics type.
// ftp should be set to the rider's known Functional Threshold Power.
// If ftp is not known, then 0 can be passed and FTP will be calculated based
// on the best 20min effort in the provided power_each_second.
func NewPowerMetrics(ftp int, power_each_second []int) PowerMetrics {
	var pm PowerMetrics

	if ftp <= 0 {
		pm.FTP = FunctionalThreshold(&power_each_second)
	} else {
		pm.FTP = ftp
	}
	pm.Time = RideTime(&power_each_second)
	pm.AP = Average(&power_each_second)
	pm.NP = Normalized(&power_each_second)
	pm.VI = VariabilityIndex(pm.NP, pm.AP)
	pm.INF = IntensityFactor(pm.NP, pm.FTP)
	pm.TSS = TrainingStressScore(pm.Time, pm.NP, pm.FTP, pm.INF)

	return pm
}
