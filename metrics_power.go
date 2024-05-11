package cycling

import (
	"math"
)

type Session struct {
	Time         int // in seconds
	PowerEachSec []int
	FTP          int
	AP           int
	NP           int
	IF           float64
	TSS          float64
}

func (s *Session) SessionTime() {
	s.Time = len(s.PowerEachSec)
}

// FunctionalThresholdPower is the estimated maximum avg power achievable for 1 hr.
// This is a rough analogy to Lactate Threshold that can be measured outside of a lab.
// Standard FTP calculation is (avg power of a 20min max-effort session) * 0.95.
func (s *Session) FunctionalThresholdPower() {
	s.FTP = int(largestSubsetAvg(s.PowerEachSec, 1200) * 0.95)
}

// NormalizedPower is a weighted average of power for a Session.
// It places more weight on higher power efforts.
// Steps to calculate are:
// --Calculate 30 second moving average power for workout.
// --Raise the reuslting values to the forth power.
// --Determine the average of those ^4 raised values.
// --Calculate 4th root of that average.
func (s *Session) NormalizedPower() {
	rolling_avgs := movingAverageInts(s.PowerEachSec, 30)
	var raised_avgs []int
	for _, v := range rolling_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	s.NP = int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

// IntensityFactor is the ratio of NormalizedPower to FunctionalThresholdPower.
// Used to determine how difficult a session was relative to a rider's capability.
func (s *Session) IntensityFactor() {
	s.IF = float64(s.NP) / float64(s.FTP)
}

// TrainingStressScore is a measurement of how taxing a session was.
// This factors in the length of the session as well as the intensity.
// The formula is (Time*NP*IF) / (FTP * 3600) * 100
func (s *Session) TrainingStressScore() {
	effort_given := float64(s.Time) * float64(s.NP) * s.IF
	baseline_effort := float64(s.FTP) * 3600
	s.TSS = effort_given / baseline_effort * 100.00
}

// /////////////////////////////////////////////////////////////////////////////
// Helper functions
func movingAverageInts(ints []int, period int) []int {
	var avgs []int
	for k := range ints {
		start := k - period
		if start < 0 {
			start = 0
		}
		period_sum := sumInts(ints[start:k])
		period_avg := math.Round(float64(period_sum) / float64(period))
		avgs = append(avgs, int(period_avg))
	}
	return avgs
}

// largestSubsetAvg finds the average of the largest sub-slice of size in slice values.
func largestSubsetAvg(ints []int, size int) float64 {
	return minAvgInts(largestSub(ints, size), size)
}

// largestSub returns the largest sub-slice of len size in the provided int slice.
func largestSub(ints []int, size int) (largest []int) {
	var sum int
	for sub_end := range ints {
		sub_start := sub_end - size
		if sub_start < 0 {
			sub_start = 0
		}
		sub := ints[sub_start:sub_end]
		sub_sum := sumInts(sub)
		if sub_sum > sum {
			sum = sub_sum
			largest = sub
		}
	}
	return largest
}

func minAvgInts(ints []int, min int) float64 {
	if len(ints) > min {
		min = len(ints)
	}
	if min < 1 {
		return 0.0
	}
	return float64(sumInts(ints)) / float64(min)
}

func avgInts(ints []int) float64 {
	if len(ints) < 1 {
		return 0.0
	}
	return float64(sumInts(ints)) / float64(len(ints))
}

func sumInts(ints []int) (sum int) {
	for _, v := range ints {
		sum = sum + v
	}
	return sum
}
