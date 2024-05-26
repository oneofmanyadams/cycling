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
	INF              float64 // Intensity Factor
	TSS              float64 // Training Stress Score
	HeartRateEachSec []int
}

// NewHeartRateMetrics is the prefered method of creating a HeartRateMetrics type.
// fthr should be set to the rider's known Functional Threshold Heart Rate.
// If fthr is not known, then 0 can be passed and FTHR will be calculated based on
// the best 20min effort in the provided hr_each_second.
func NewHeartRateMetrics(fthr int, hr_each_second []int) HeartRateMetrics {
	var hrm HeartRateMetrics

	if fthr <= 0 {
		hrm.FTHR = hrm.FunctionalThresholdHeartRate(&hr_each_second)
	} else {
		hrm.FTHR = fthr
	}
	hrm.HeartRateEachSec = hr_each_second
	hrm.Time = hrm.SessionTime(&hr_each_second)
	hrm.AHR = hrm.AverageHeartRate(&hr_each_second)
	hrm.NHR = hrm.NormalizedHeartRate(&hr_each_second)
	hrm.VI = hrm.VariabilityIndex(hrm.NHR, hrm.AHR)
	hrm.INF = hrm.IntensityFactor(hrm.NHR, hrm.FTHR)
	hrm.TSS = hrm.TrainingStressScore(hrm.Time, hrm.NHR, hrm.FTHR, hrm.INF)

	return hrm
}

// SessionTime calculates Time based on total number of elements in HeartRateEachSec.
func (s *HeartRateMetrics) SessionTime(hr_each_sec *[]int) int {
	return len(*hr_each_sec)
}

// FunctionalThresholdHeartRate is just a duplication of the formula to
// calculate FTP, see power metric FunctionalThresholdPower for more info.
func (s *HeartRateMetrics) FunctionalThresholdHeartRate(hr_each_sec *[]int) int {
	return int(largestSubsetAvg(*hr_each_sec, 1200) * 0.95)
}

// AverageHeartRate calculates the average heart rate for a session.
func (s *HeartRateMetrics) AverageHeartRate(hr_each_sec *[]int) int {
	return int(avgInts(*hr_each_sec))
}

// NormalizedHeartRate is a weighted average of HeartRateEachSec, intended to
// give more weight to higher intensity efforts.
// Details on calculation can be found on the NormalizedPower method of PowerMetrics.
func (s *HeartRateMetrics) NormalizedHeartRate(hr_each_sec *[]int) int {
	moving_avgs := movingAverageInts(*hr_each_sec, 30)
	var raised_avgs []int
	for _, v := range moving_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	return int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

// VariabilityIndex is the ratio of NormalizedHeartRate to AverageHeartRate.
// A number close to 1 means that HeartRate did not fluctuate much.
func (s *HeartRateMetrics) VariabilityIndex(nhr, ahr int) float64 {
	return float64(nhr) / float64(ahr)
}

// IntensityFactor is the ratio of NormalizedHeartRate to FunctionalThresholdHeartRate.
// The larger the number, the harder the session was.
// An IF of 1 basically means a session was done right at Threshold effort.
func (s *HeartRateMetrics) IntensityFactor(nhr, fthr int) float64 {
	return float64(nhr) / float64(fthr)
}

// TrainingStressScore measures how difficult a session was relative
// to an individual's FunctionalThresholdHeartRate.
// Duplicates the same formula from PowerMetric's method of the same name.
func (s *HeartRateMetrics) TrainingStressScore(time, nhr, fthr int, inf float64) float64 {
	effort_given := float64(time) * float64(nhr) * inf
	baselines_effort := float64(fthr) * 3600
	return effort_given / baselines_effort * 100.00
}
