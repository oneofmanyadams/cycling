package cycling

type Trends struct {
	CTL float64 // Cronic Training Load
	ATL float64 // Acute Training Load
	TSB float64 // Training Stress Balance
}

// CronicTrainingLoad is a measure of current fitness and is an
// average of the past 42 days of training strss scores.
func CronicTrainingLoad(stress_scores []float64) float64 {
	if len(stress_scores) > 42 {
		stress_scores = stress_scores[len(stress_scores)-42:]
	}
	return minAvgFloats(stress_scores, 42)
}

// AcuteTrainingLoad gauges how much workload is being put on an individual based
// on their average training stress score over the past 7 days.
func AcuteTrainingLoad(stress_scores []float64) float64 {
	if len(stress_scores) > 7 {
		stress_scores = stress_scores[len(stress_scores)-7:]
	}
	return minAvgFloats(stress_scores, 7)
}

// TrainingStressBalance compares recent avg workload vs longer term avg workloads.
// A negative number indicates training above average historic workloads.
func TrainingStressBalance(ctl float64, atl float64) float64 {
	return ctl - atl
}
