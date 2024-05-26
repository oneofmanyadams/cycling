package cycling

import "math"

// RideTime calculates the total ride time in seconds based on a
// provided Ride metric (like PowerEachSecond).
func RideTime(ride_metric *[]int) int {
	return len(*ride_metric)
}

// FunctionalThreshold can be used to calculate FTP for FTHR based
// on the popular FTP formula of: "avg power of a 20min max effort" * 95%.
func FunctionalThreshold(ride_metric *[]int) int {
	return int(largestSubsetAvg(*ride_metric, 1200) * 0.95)
}

// Average does exactly what you would think for metrics like Power or HR.
func Average(ride_metric *[]int) int {
	return int(avgInts(*ride_metric))
}

// Normalized is based on the formula for determining Normalized Power.
// The reason for this metric is more heavily weigh high intensity efforts.
// The calculation is:
// -Calculate 30 second moving average of power for a given ride.
// -Raise the resulting values to the fourth power.
// -Determine the average of those ^4 raised values.
// -Calculate the 4th root of that average.
func Normalized(ride_metric *[]int) int {
	moving_avgs := movingAverageInts(*ride_metric, 30)
	var raised_avgs []int
	for _, v := range moving_avgs {
		raised_avgs = append(raised_avgs, int(math.Pow(float64(v), 4)))
	}
	avg_of_raised := avgInts(raised_avgs)
	return int(math.Round(math.Pow(avg_of_raised, 1.0/4.0)))
}

// VariabilityIndex calculates the ratio of a metrics normalized value
// vs it's average value. A number close to 1 means the metric did not
// fluctuate much over the course of a ride.
func VariabilityIndex(normalized, average int) float64 {
	return float64(normalized) / float64(average)
}

// IntensityFactor is the ratio of a metrics normalized value vs it's
// funtional threshold. The larger the number the harder the ride was.
// A value of 1 essentially means the ride was done right at threshold.
func IntensityFactor(normalized, threshold int) float64 {
	return float64(normalized) / float64(threshold)
}

// TrainingStressScore measures how difficult a ride was relative to an
// individual's Funtional Threshold. This factors in the duration and intensity.
// The formula (based on power) is: (Time * NP * IF) / (FTP * 3600) *100.
func TrainingStressScore(time, normalized, functional int, inf float64) float64 {
	effort_given := float64(time) * float64(normalized) * inf
	baseline_effort := float64(functional) * 3600
	return effort_given / baseline_effort * 100.00
}
