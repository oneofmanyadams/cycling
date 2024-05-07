package cycling

type Session struct {
	Time         int // in seconds
	PowerEachSec []int
	Ftp          int
}

type TrainingStressScore struct {
	Time int
	NP   NormalizedPower
	IP   IntensityFactor
	FTP  FunctionalThresholdPower
	// (Time*NP*IF) / (FTP * 3600) * 100
}

type NormalizedPower struct {
	PowerEachSecond []int
	// Calculate rolling 30 second average power for workout.
	// Raise resulting values to forth power
	// Determine average of the ^4 values
	// Determine 4th root of that average
}

type IntensityFactor struct {
	NP  NormalizedPower
	FTP FunctionalThresholdPower
	// NP / FTP
}

type FunctionalThresholdPower struct {
	// Maximum average power for 1hr
	// Or Maximum average power for 20mins * .95
}
