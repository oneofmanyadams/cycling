package cycling

import (
	"errors"
	"fmt"
)

type PowerZones []PowerZone

func (s *PowerZones) MatchZone(pwr int) (PowerZone, error) {
	for _, zone := range *s {
		if zone.Matches(pwr) {
			return zone, nil
		}
	}
	return PowerZone{}, errors.New(fmt.Sprintf("No PowerZone for %d", pwr))
}

type PowerZone struct {
	Number      int
	Name        string
	Description string
	MinWatts    int
	MaxWatts    int
}

func (s *PowerZone) Matches(power int) bool {
	return (s.MinWatts <= power) && (s.MaxWatts >= power)
}

func (s *PowerZone) AvgWatts() int {
	return (s.MaxWatts + s.MinWatts) / 2
}
