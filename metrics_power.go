package cycling

import "math"

// PowerMetrics aggregates and calculates various Power Metrics.
// The minimum input neededis PowerEachSec, all other fields can be
// populated based on that data.
// The function NewPowerMetrics() is the prefered way to generate this type.
type PowerMetrics struct {
	Time         int     // in seconds
	FTP          int     // Functional Threshold Power
	AP           int     // Average Power
	NP           int     // Normalized Power
	VI           float64 // Variability Index
	INF          float64 // Intensity Factor
	TSS          float64 // Training Stress Score
	PowerEachSec []int
}

// NewPowerMetrics is the prefered method of creating a PowerMetrics type.
// ftp should be set to the rider's known Functional Threshold Power.
// If ftp is not known, then 0 can be passed and FTP will be calculated based
// on the best 20min effort in the provided power_each_second.
func NewPowerMetrics(ftp int, power_each_second []int) PowerMetrics {
	var pm PowerMetrics

	if ftp <= 0 {
		pm.FTP = pm.FunctionalThresholdPower(&power_each_second)
	} else {
		pm.FTP = ftp
	}
	pm.PowerEachSec = power_each_second
	pm.Time = pm.SessionTime(&power_each_second)
	pm.AP = pm.AveragePower(&power_each_second)
	pm.NP = pm.NormalizedPower(&power_each_second)
	pm.VI = pm.VariabilityIndex(pm.NP, pm.AP)
	pm.INF = pm.IntensityFactor(pm.NP, pm.FTP)
	pm.TSS = pm.TrainingStressScore(pm.Time, pm.NP, pm.FTP, pm.INF)

	return pm
}

// SessionTime calculates Time based on total number of elements in PowerEachSec.
func (s *PowerMetrics) SessionTime(power_each_sec *[]int) int {
	return len(*power_each_sec)
}

// FunctionalThresholdPower is the estimated maximum avg power achievable for 1 hr.
// This is a rough analogy to Lactate Threshold that can be measured outside of a lab.
// Standard FTP calculation is (avg power of a 20min max-effort session) * 0.95.
func (s *PowerMetrics) FunctionalThresholdPower(pow_each_sec *[]int) int {
	return int(largestSubsetAvg(*pow_each_sec, 1200) * 0.95)
}

// AveragePower calculates the average power for a session.
func (s *PowerMetrics) AveragePower(pow_each_sec *[]int) int {
	return int(avgInts(*pow_each_sec))
}

// NormalizedPower is a weighted average of PowerEachSec, intended to
// give more weight to higher intensity efforts.
// Steps to calculate are:
// --Calculate 30 second moving average power for workout.
// --Raise the reuslting values to the forth power.
// --Determine the average of those ^4 raised values.
// --Calculate 4th root of that average.
func (s *PowerMetrics) NormalizedPower(pow_each_sec *[]int) int {
	rolling_avgs := movingAverageInts(*pow_each_sec, 30)
	var raised_avgs []int
	for _, v := range rolling_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	return int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

// VariabilityIndex is the ratio of NormalizedPower to AveragePower.
// A number close to 1 means that Power did not fluctuate much.
func (s *PowerMetrics) VariabilityIndex(np, ap int) float64 {
	return float64(np) / float64(ap)
}

// IntensityFactor is the ratio of NormalizedPower to FunctionalThresholdPower.
// The larger the number, the harder the session was.
// An IF of 1 basically means a session was done right at Threshold effort.
func (s *PowerMetrics) IntensityFactor(np, ftp int) float64 {
	return float64(np) / float64(ftp)
}

// TrainingStressScore measures how difficult a session was relative
// to an individual's FunctionalThresholdPower.
// This factors in the length of the session as well as the intensity.
// The formula is (Time*NP*IF) / (FTP * 3600) * 100
func (s *PowerMetrics) TrainingStressScore(time, np, ftp int, inf float64) float64 {
	effort_given := float64(time) * float64(np) * inf
	baseline_effort := float64(ftp) * 3600
	return effort_given / baseline_effort * 100.00
}
