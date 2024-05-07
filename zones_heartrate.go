package cycling

type HeartZone struct {
	Number      int
	Name        string
	Description string
	MinBPM      int
	MaxBPM      int
}

func (s *HeartZone) AvgBPM() int {
	return (s.MaxBPM + s.MinBPM) / 2
}
