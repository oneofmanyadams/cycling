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
	IF           float64 // Intensity Factor
	TSS          float64 // Training Stress Score
	PowerEachSec []int
}

// NewPowerMetrics is the prefered method of creating a PowerMetrics type.
// ftp should be set to the rider's known Functional Threshold Power.
// If ftp is not known, then 0 can be passed and FTP will be calculated based
// on the best 20min effort in the provided power_each_second.
func NewPowerMetrics(ftp int, power_each_second []int) PowerMetrics {
	var pm PowerMetrics
	pm.FTP = ftp
	pm.PowerEachSec = power_each_second
	pm.CalculateMetrics()
	return pm
}

// CalculateMetrics runs all metrics calculating methods in the correct
// order they need to be called in to corrrectly populate PowerMetrics.
// FunctionalThresholdPower is not called if it is already set (>0).
func (s *PowerMetrics) CalculateMetrics() {
	s.SessionTime()
	if s.FTP <= 0 {
		s.FunctionalThresholdPower()
	}
	s.AveragePower()
	s.NormalizedPower()
	s.VariabilityIndex()
	s.IntensityFactor()
	s.TrainingStressScore()
}

// SessionTime calculates Time based on total number of elements in PowerEachSec.
func (s *PowerMetrics) SessionTime() {
	s.Time = len(s.PowerEachSec)
}

// FunctionalThresholdPower is the estimated maximum avg power achievable for 1 hr.
// This is a rough analogy to Lactate Threshold that can be measured outside of a lab.
// Standard FTP calculation is (avg power of a 20min max-effort session) * 0.95.
func (s *PowerMetrics) FunctionalThresholdPower() {
	s.FTP = int(largestSubsetAvg(s.PowerEachSec, 1200) * 0.95)
}

// AveragePower calculates the average power for a session.
func (s *PowerMetrics) AveragePower() {
	s.AP = int(avgInts(s.PowerEachSec))
}

// NormalizedPower is a weighted average of PowerEachSec, intended to
// give more weight to higher intensity efforts.
// Steps to calculate are:
// --Calculate 30 second moving average power for workout.
// --Raise the reuslting values to the forth power.
// --Determine the average of those ^4 raised values.
// --Calculate 4th root of that average.
func (s *PowerMetrics) NormalizedPower() {
	rolling_avgs := movingAverageInts(s.PowerEachSec, 30)
	var raised_avgs []int
	for _, v := range rolling_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	s.NP = int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

// VariabilityIndex is the ratio of NormalizedPower to AveragePower.
// A number close to 1 means that Power did not fluctuate much.
func (s *PowerMetrics) VariabilityIndex() {
	s.VI = float64(s.NP) / float64(s.AP)
}

// IntensityFactor is the ratio of NormalizedPower to FunctionalThresholdPower.
// The larger the number, the harder the session was.
// An IF of 1 basically means a session was done right at Threshold effort.
func (s *PowerMetrics) IntensityFactor() {
	s.IF = float64(s.NP) / float64(s.FTP)
}

// TrainingStressScore measures how difficult a session was relative
// to an individual's FunctionalThresholdPower.
// This factors in the length of the session as well as the intensity.
// The formula is (Time*NP*IF) / (FTP * 3600) * 100
func (s *PowerMetrics) TrainingStressScore() {
	effort_given := float64(s.Time) * float64(s.NP) * s.IF
	baseline_effort := float64(s.FTP) * 3600
	s.TSS = effort_given / baseline_effort * 100.00
}
