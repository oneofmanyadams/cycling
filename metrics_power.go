package cycling

// --ToDo
// -- Finish NormalizedPower function
// -- Test with real data and compare to app results.

import (
	"fmt"
	"math"
)

type Session struct {
	Time         int // in seconds
	PowerEachSec []int
	FTP          int
	NP           int
	IF           float64
	TSS          float64
}

// FunctionalThresholdPower is the estimated maximum avg power achievable for 1 hr.
// This is a rough analogy to Lactate Threshold that can be measured outside of a lab.
// Standard FTP calculation is avg power of a 20min max-effort session * 0.95.
func (s *Session) FunctionalThresholdPower() int {
	if s.FTP == 0 {
		s.FTP = int(findHighestChunkAverage(1200, s.PowerEachSec) * 0.95)
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
		// this won't work, numbers are too big.
		// likely need to sum using math/big package and then implement
		// own 4 root function.
		// Try using int64 first? or even uint64?
		avg_of_raised := avgIntSlice(raised_avgs)
		s.NP = int(math.Pow(avg_of_raised, 1/4))
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
		effort_given := float64(s.Time) *
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
		chunk_sum := sumIntSlice(values[chunk_start:k])
		chunk_average := math.Round(float64(chunk_sum) / float64(chunk_len))
		averages_slice = append(averages_slice, int(chunk_average))
	}
	return averages_slice
}

func findHighestChunkAverage(chunk_len int, values []int) float64 {
	var largest float64
	for k := range values {
		var calc_val float64
		// Is this actually a bug?
		// For slices with len less than chunk_len should
		// we still calc the average based on chunk_len?
		// (instead of sub slice len which is what this is doing?)
		if k < chunk_len {
			calc_val = avgIntSlice(values[0:k])
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
