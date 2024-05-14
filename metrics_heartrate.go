package cycling

import "math"

// HeartRateMetrics aggregates and calculates various Heart Rate Metrics.
// The minimum input needed is HeartRateEachSec, all other fields can be
// populated based on that data.
// The function NewHeartRateMetrics() is the prefered way to generate this type.
type HeartRateMetrics struct {
	Time             int     // in seconds
	FTHR             int     // Funtional Threshold HeartRate
	AHR              int     // Average Heart Rate
	NHR              int     // Normalized Heart Rate
	VI               float64 // Variability Index
	IF               float64 // Intensity Factor
	TSS              float64 // Training Stress Score
	HeartRateEachSec []int
}

// NewHeartRateMetrics is the prefered method of creating a HeartRateMetrics type.
// fthr should be set to the rider's known Functional Threshold Heart Rate.
// If fthr is not known, then 0 can be passed and FTHR will be calculated based on
// the best 20min effort in the provided hr_each_second.
func NewHeartRateMetrics(fthr int, hr_each_second []int) HeartRateMetrics {
	var hrm HeartRateMetrics
	hrm.FTHR = fthr
	hrm.HeartRateEachSec = hr_each_second
	hrm.CalculateMetrics()
	return hrm
}

// CalculateMetrics runs all metrics caluclating methods in the correct
// order they need to be called in to correctly populate a HeartRateMetrics.
// FunctionalThresholdHeartRate is not called if it is already set (>0).
func (s *HeartRateMetrics) CalculateMetrics() {
	s.SessionTime()
	if s.FTHR <= 0 {
		s.FunctionalThresholdHeartRate()
	}
	s.AverageHeartRate()
	s.NormalizedHeartRate()
	s.VariabilityIndex()
	s.IntensityFactor()
	s.TrainingStressScore()
}

// SessionTime calculates Time based on total number of elements in HeartRateEachSec.
func (s *HeartRateMetrics) SessionTime() {
	s.Time = len(s.HeartRateEachSec)
}

// FunctionalThresholdHeartRate is just a duplication of the formula to
// calculate FTP, see power metric FunctionalThresholdPower for more info.
func (s *HeartRateMetrics) FunctionalThresholdHeartRate() {
	s.FTHR = int(largestSubsetAvg(s.HeartRateEachSec, 1200) * 0.95)
}

// AverageHeartRate calculates the average heart rate for a session.
func (s *HeartRateMetrics) AverageHeartRate() {
	s.AHR = int(avgInts(s.HeartRateEachSec))
}

// NormalizedHeartRate is a weighted average of HeartRateEachSec, intended to
// give more weight to higher intensity efforts.
// Details on calculation can be found on the NormalizedPower method of PowerMetrics.
func (s *HeartRateMetrics) NormalizedHeartRate() {
	moving_avgs := movingAverageInts(s.HeartRateEachSec, 30)
	var raised_avgs []int
	for _, v := range moving_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	s.NHR = int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

// VariabilityIndex is the ratio of NormalizedHeartRate to AverageHeartRate.
// A number close to 1 means that HeartRate did not fluctuate much.
func (s *HeartRateMetrics) VariabilityIndex() {
	s.VI = float64(s.NHR) / float64(s.AHR)
}

// IntensityFactor is the ratio of NormalizedHeartRate to FunctionalThresholdHeartRate.
// The larger the number, the harder the session was.
// An IF of 1 basically means a session was done right at Threshold effort.
func (s *HeartRateMetrics) IntensityFactor() {
	s.IF = float64(s.NHR) / float64(s.FTHR)
}

// TrainingStressScore measures how difficult a session was relative
// to an individual's FunctionalThresholdHeartRate.
// Duplicates the same formula from PowerMetric's method of the same name.
func (s *HeartRateMetrics) TrainingStressScore() {
	effort_given := float64(s.Time) * float64(s.NHR) * s.IF
	baselines_effort := float64(s.FTHR) * 3600
	s.TSS = effort_given / baselines_effort * 100.00
}
