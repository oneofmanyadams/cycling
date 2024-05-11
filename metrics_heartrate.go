package cycling

type HearrRateMetrics struct {
	HeartRateEachSec []int
	Time             int     // in seconds
	FTHR             int     // FuntionalThresholdHeartRate
	AHR              int     // AverageHeartRate
	NHR              int     // NormalizedHeartRate
	VI               float64 // VariabilityIndex
	IF               float64 // IntensityFactor
	TSS              float64 // TrainingStressScore
}
