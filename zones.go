package cycling

type Zones []Zone

// Load from template functions

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

// Max []float64 methods
func (s *Zones) MaxHearRates() []float64 {
	var heart_rates []float64
	for _, z := range *s {
		heart_rates = append(heart_rates, z.MaxHeartRate)
	}
	return heart_rates
}

func (s *Zones) MaxPowers() []float64 {
	var powers []float64
	for _, z := range *s {
		powers = append(powers, z.MaxPower)
	}
	return powers
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
