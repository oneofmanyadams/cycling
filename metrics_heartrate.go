package cycling

// HeartRateMetrics aggregates and calculates various Heart Rate Metrics.
// The minimum input needed is HeartRateEachSec, all other fields can be
// populated based on that data.
// The function NewHeartRateMetrics() is the prefered way to generate this type.
type HeartRateMetrics struct {
	Time int     // in seconds
	FTHR int     // Funtional Threshold HeartRate
	AHR  int     // Average Heart Rate
	NHR  int     // Normalized Heart Rate
	VI   float64 // Variability Index
	INF  float64 // Intensity Factor
	TSS  float64 // Training Stress Score
}

// NewHeartRateMetrics is the prefered method of creating a HeartRateMetrics type.
// fthr should be set to the rider's known Functional Threshold Heart Rate.
// If fthr is not known, then 0 can be passed and FTHR will be calculated based on
// the best 20min effort in the provided hr_each_second.
func NewHeartRateMetrics(fthr int, hr_each_second []int) HeartRateMetrics {
	var hrm HeartRateMetrics

	if fthr <= 0 {
		hrm.FTHR = FunctionalThreshold(&hr_each_second)
	} else {
		hrm.FTHR = fthr
	}
	hrm.Time = RideTime(&hr_each_second)
	hrm.AHR = Average(&hr_each_second)
	hrm.NHR = Normalized(&hr_each_second)
	hrm.VI = VariabilityIndex(hrm.NHR, hrm.AHR)
	hrm.INF = IntensityFactor(hrm.NHR, hrm.FTHR)
	hrm.TSS = TrainingStressScore(hrm.Time, hrm.NHR, hrm.FTHR, hrm.INF)

	return hrm
}
