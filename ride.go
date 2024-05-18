package cycling

type Ride struct {
	TimeStampEachSec   []int64
	PowerEachSec       []int
	HeartRateEachSec   []int
	CadenceEachSec     []int
	SpeedEachSec       []float64
	DistanceCumulative []float64
	CaloriesCumulative []int
}

func (s *Ride) AddTime(t int64) {
	s.TimeStampEachSec = append(s.TimeStampEachSec, t)
}
func (s *Ride) AddPower(p int) {
	s.PowerEachSec = append(s.PowerEachSec, p)
}
func (s *Ride) AddHeartRate(hr int) {
	s.HeartRateEachSec = append(s.HeartRateEachSec, hr)
}
func (s *Ride) AddCadence(c int) {
	s.CadenceEachSec = append(s.CadenceEachSec, c)
}
func (s *Ride) AddSpeed(sp float64) {
	s.SpeedEachSec = append(s.SpeedEachSec, sp)
}
func (s *Ride) AddDistance(d float64) {
	s.DistanceCumulative = append(s.DistanceCumulative, d)
}
func (s *Ride) AddCalories(cal int) {
	s.CaloriesCumulative = append(s.CaloriesCumulative, cal)
}
