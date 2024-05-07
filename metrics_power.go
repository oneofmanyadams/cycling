package cycling

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
		return s.NP
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
	if len(values) < chunk_len {
		return []int{}
	}
	// rolling average logic here.
	return []int{}
}

func findHighestChunkAverage(chunk_len int, values []int) float64 {
	var largest float64
	for k := range values {
		var calc_val float64
		if k < chunk_len {
			calc_val = avgIntSlice(values[0:k])
		} else {
			calc_val = avgIntSlice(values[k-chunk_len : k])
		}
		if calc_val > largest {
			largest = calc_val
		}
	}
	// Do a rolling sum of current val plus prev chunklen_records.
	// Record whatever the largest value is.
	// return largest_val/chunk_len

	return largest
}

func avgIntSlice(values []int) float64 {
	var sum int
	for _, v := range values {
		sum = sum + v
	}
	return float64(sum) / float64(len(values))
}
