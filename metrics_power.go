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

///////////////////////////////////////////////////////////////////////////////
// Helper functions

// largestSubsetAvg finds the average of the largest sub-slice of size chunk_len in slice values.
func largestSubsetAvg(values []int, size int) float64 {
	return avgInts(largestSub(values, size), size)
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

// avgInts allows a min_denominator, which makes it a little easier to do moving avgs
func avgInts(ints []int, min_denominator int) float64 {
	if len(ints) > min_denominator {
		min_denominator = len(ints)
	}
	if min_denominator == 0 {
		return 0.0
	}
	return float64(sumInts(ints)) / float64(min_denominator)
}

func sumInts(ints []int) (sum int) {
	for _, v := range ints {
		sum = sum + v
	}
	return sum
}

///////////////////////////////////////////////////////////////////////////////
// Deprecated functions below here

func (s *Session) CalculatedTime() int {
	if s.Time == 0 {
		s.Time = len(s.PowerEachSec)
	}
	return s.Time
}

// FunctionalThresholdPower is the estimated maximum avg power achievable for 1 hr.
// This is a rough analogy to Lactate Threshold that can be measured outside of a lab.
// Standard FTP calculation is avg power of a 20min max-effort session * 0.95.
func (s *Session) FunctionalThresholdPower() int {
	if s.FTP == 0 {
		s.FTP = int(largestSubsetAvg(s.PowerEachSec, 1200) * 0.95)
	}
	return s.FTP
}

// NormalizedPower is a rolling weighted average for power for a Session.
// It essentially gives more credit for efforts done above FTP.
// Steps to calculate are:
// --Calculate rolling 30 second average power for workout.
// --Raise the reuslting values to the forth power.
// --Determine the average of those ^4 raised values.
// --Calculate 4th root of that average.
func (s *Session) NormalizedPower() int {
	if s.NP == 0 {
		rolling_avgs := calcRollingAverage(30, s.PowerEachSec)
		var raised_avgs []int
		for _, v := range rolling_avgs {
			raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
		}
		avg_of_raised := avgInts(raised_avgs, 0)
		s.NP = int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
	}
	return s.NP
}

// IntensityFactor is the ratio of NormalizedPower to FunctionalThresholdPower.
// Used to determine how difficult a session was relative to a rider's capability.
func (s *Session) IntensityFactor() float64 {
	if s.IF == 0 {
		s.IF = float64(s.NormalizedPower()) /
			float64(s.FunctionalThresholdPower())
	}
	return s.IF
}

// TrainingStressScore is a measurement of how taxing a session was.
// This factors in the length of the session as well as the intensity.
// The formula is (Time*NP*IF) / (FTP * 3600) * 100
func (s *Session) TrainingStressScore() float64 {
	if s.TSS == 0 {
		effort_given := float64(s.CalculatedTime()) *
			float64(s.NormalizedPower()) *
			s.IntensityFactor()
		baseline_effort := float64(s.FunctionalThresholdPower()) * 3600
		s.TSS = effort_given / baseline_effort * 100.00
	}
	return s.TSS
}

// Helper functions
func calcRollingAverage(chunk_len int, values []int) []int {
	var averages_slice []int
	for k := range values {
		chunk_start := k - chunk_len
		if k < chunk_len {
			chunk_start = 0
		}
		chunk_sum := sumInts(values[chunk_start:k])
		chunk_average := math.Round(float64(chunk_sum) / float64(chunk_len))
		averages_slice = append(averages_slice, int(chunk_average))
	}
	return averages_slice
}

// DONE
/*
func findHighestChunkAverage(chunk_len int, values []int) float64 {
	var largest float64
	for k := range values {
		var calc_val float64
		// Is this actually a bug?
		// For slices with len less than chunk_len should
		// we still calc the average based on chunk_len?
		// (instead of sub slice len which is what this is doing?)
		if k < chunk_len {
			calc_val = float64(sumIntSlice(values[0:k])) / float64(chunk_len)
		} else {
			calc_val = avgIntSlice(values[k-chunk_len : k])
		}
		if calc_val > largest {
			largest = calc_val
		}
	}
	return largest
}

func avgIntSlice(values []int) float64 {
	return float64(sumIntSlice(values)) / float64(len(values))
}

func sumIntSlice(values []int) int {
	var sum int
	for _, v := range values {
		sum = sum + v
	}
	return sum
}
*/
