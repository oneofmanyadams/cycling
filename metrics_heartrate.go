package cycling

import "math"

type HeartRateMetrics struct {
	HeartRateEachSec []int
	Time             int     // in seconds
	FTHR             int     // FuntionalThresholdHeartRate
	AHR              int     // AverageHeartRate
	NHR              int     // NormalizedHeartRate
	VI               float64 // VariabilityIndex
	IF               float64 // IntensityFactor
	TSS              float64 // TrainingStressScore
}

func NewHeartRateMetrics(fthr int, hr_each_second []int) HeartRateMetrics {
	var hrm HeartRateMetrics
	hrm.FTHR = fthr
	hrm.HeartRateEachSec = hr_each_second
	hrm.CalculateMetrics()
	return hrm
}

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

func (s *HeartRateMetrics) SessionTime() {
	s.Time = len(s.HeartRateEachSec)
}

// FunctionalThresholdHeartRate is just a duplication of the formula to
// calculate FTP, see power metric FunctionalThresholdPower for more info.
func (s *HeartRateMetrics) FunctionalThresholdHeartRate() {
	s.FTHR = int(largestSubsetAvg(s.HeartRateEachSec, 1200) * 0.95)
}

func (s *HeartRateMetrics) AverageHeartRate() {
	s.AHR = int(avgInts(s.HeartRateEachSec))
}

func (s *HeartRateMetrics) NormalizedHeartRate() {
	moving_avgs := movingAverageInts(s.HeartRateEachSec, 30)
	var raised_avgs []int
	for _, v := range moving_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	s.NHR = int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

func (s *HeartRateMetrics) VariabilityIndex() {
	s.VI = float64(s.NHR) / float64(s.AHR)
}

func (s *HeartRateMetrics) IntensityFactor() {
	s.IF = float64(s.NHR) / float64(s.FTHR)
}

func (s *HeartRateMetrics) TrainingStressScore() {
	effort_given := float64(s.Time) * float64(s.NHR) * s.IF
	baselines_effort := float64(s.FTHR) * 3600
	s.TSS = effort_given / baselines_effort * 100.00
}
