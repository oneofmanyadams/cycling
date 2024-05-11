package cycling

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
