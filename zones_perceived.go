package cycling

type PerceivedZone struct {
	Number      int
	Name        string
	Description string
	MinExertion int
	MaxExertion int
}

func (s *PerceivedZone) AvgExertion() int {
	return (s.MaxExertion + s.MinExertion) / 2
}
