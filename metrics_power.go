package cycling

import (
	"math"
)

type PowerMetrics struct {
	PowerEachSec []int
	Time         int // in seconds
	FTP          int
	AP           int
	NP           int
	VI           float64
	IF           float64
	TSS          float64
}

func NewPowerMetrics(ftp int, power_each_second []int) PowerMetrics {
	var pm PowerMetrics
	pm.FTP = ftp
	pm.PowerEachSec = power_each_second
	pm.CalculateMetrics()
	return pm
}

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

func (s *PowerMetrics) SessionTime() {
	s.Time = len(s.PowerEachSec)
}

// FunctionalThresholdPower is the estimated maximum avg power achievable for 1 hr.
// This is a rough analogy to Lactate Threshold that can be measured outside of a lab.
// Standard FTP calculation is (avg power of a 20min max-effort session) * 0.95.
func (s *PowerMetrics) FunctionalThresholdPower() {
	s.FTP = int(largestSubsetAvg(s.PowerEachSec, 1200) * 0.95)
}

func (s *PowerMetrics) AveragePower() {
	s.AP = int(avgInts(s.PowerEachSec))
}

// NormalizedPower is a weighted average of power for a Session.
// It places more weight on higher power efforts.
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

// VariabilityIndex measure how smooth power output was during a ride.
func (s *PowerMetrics) VariabilityIndex() {
	s.VI = float64(s.NP) / float64(s.AP)
}

// IntensityFactor is the ratio of NormalizedPower to FunctionalThresholdPower.
// Used to determine how difficult a session was relative to a rider's capability.
func (s *PowerMetrics) IntensityFactor() {
	s.IF = float64(s.NP) / float64(s.FTP)
}

// TrainingStressScore is a measurement of how taxing a session was.
// This factors in the length of the session as well as the intensity.
// The formula is (Time*NP*IF) / (FTP * 3600) * 100
func (s *PowerMetrics) TrainingStressScore() {
	effort_given := float64(s.Time) * float64(s.NP) * s.IF
	baseline_effort := float64(s.FTP) * 3600
	s.TSS = effort_given / baseline_effort * 100.00
}
