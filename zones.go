package cycling

type Zones []Zone

// Functions to fulfill sort interface.
func (s Zones) Len() int {
	return len(s)
}
func (s Zones) Less(i, j int) bool {
	if len(s) <= max(i, j) {
		return false
	}
	return s[i].Level < s[j].Level
}
func (s Zones) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Zone struct {
	Level        int
	Name         string
	Description  string
	MinDuration  int
	MaxDuration  int
	MaxHeartRate float64
	MaxPower     float64
}
