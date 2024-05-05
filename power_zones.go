package cycling

type PowerZone struct {
	Number      int
	Name        string
	Description string
	MinWatts    int
	MaxWatts    int
}

func (s *PowerZone) AvgWatts() int {
	return (s.MaxWatts + s.MinWatts) / 2
}
