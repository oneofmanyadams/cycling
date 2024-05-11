package cycling

import (
	"math"
)

func movingAverageInts(ints []int, period int) []int {
	var avgs []int
	for k := range ints {
		start := k - period
		if start < 0 {
			start = 0
		}
		period_sum := sumInts(ints[start:k])
		period_avg := math.Round(float64(period_sum) / float64(period))
		avgs = append(avgs, int(period_avg))
	}
	return avgs
}

// largestSubsetAvg finds the average of the largest sub-slice of size in slice values.
func largestSubsetAvg(ints []int, size int) float64 {
	return minAvgInts(largestSub(ints, size), size)
}

// largestSub returns the largest sub-slice of len size in the provided int slice.
func largestSub(ints []int, size int) (largest []int) {
	var sum int
	for sub_end := range ints {
		sub_start := sub_end - size
		if sub_start < 0 {
			sub_start = 0
		}
		sub := ints[sub_start:sub_end]
		sub_sum := sumInts(sub)
		if sub_sum > sum {
			sum = sub_sum
			largest = sub
		}
	}
	return largest
}

func minAvgInts(ints []int, min int) float64 {
	if len(ints) > min {
		min = len(ints)
	}
	if min < 1 {
		return 0.0
	}
	return float64(sumInts(ints)) / float64(min)
}

func avgInts(ints []int) float64 {
	if len(ints) < 1 {
		return 0.0
	}
	return float64(sumInts(ints)) / float64(len(ints))
}

func sumInts(ints []int) (sum int) {
	for _, v := range ints {
		sum = sum + v
	}
	return sum
}
