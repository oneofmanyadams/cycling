package cycling

type Session struct {
	TimeStampEachSec   []int64
	PowerEachSec       []int
	HeartRateEachSec   []int
	CadenceEachSec     []int
	SpeedEachSec       []float64
	DistanceCumulative []float64
	CaloriesCumulative []int
}

func (s *Session) AddTime(t int64) {
	s.TimeStampEachSec = append(s.TimeStampEachSec, t)
}
func (s *Session) AddPower(p int) {
	s.PowerEachSec = append(s.PowerEachSec, p)
}
func (s *Session) AddHeartRate(hr int) {
	s.HeartRateEachSec = append(s.HeartRateEachSec, hr)
}
func (s *Session) AddCadence(c int) {
	s.CadenceEachSec = append(s.CadenceEachSec, c)
}
func (s *Session) AddSpeed(sp float64) {
	s.SpeedEachSec = append(s.SpeedEachSec, sp)
}
func (s *Session) AddDistance(d float64) {
	s.DistanceCumulative = append(s.DistanceCumulative, d)
}
func (s *Session) AddCalories(cal int) {
	s.CaloriesCumulative = append(s.CaloriesCumulative, cal)
}
