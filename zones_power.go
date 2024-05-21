package cycling

import (
	"errors"
	"fmt"
	"sort"
)

type PowerZones []PowerZone

// NewPowerZones calculates power zones based on a rirder's ftp as well as a
// list of of each zone's max_ftp (in terrms of a %of the ftp).
func NewPowerZones(ftp int, max_ftps []float64) PowerZones {
	sort.Float64s(max_ftps)
	var pzs PowerZones
	for len(pzs) < len(max_ftps) {
		var new_zone PowerZone
		new_zone.Number = len(pzs) + 1
		if new_zone.Number == 1 {
			new_zone.MinWatts = 0
		} else {
			new_zone.MinWatts = pzs[len(pzs)-1].MaxWatts + 1
		}
		new_zone.MaxWatts = int(float64(ftp) * max_ftps[len(pzs)])
		pzs = append(pzs, new_zone)
	}
	return pzs
}

func (s *PowerZones) MatchZone(pwr int) (PowerZone, error) {
	for _, zone := range *s {
		if zone.Matches(pwr) {
			return zone, nil
		}
	}
	return PowerZone{}, errors.New(fmt.Sprintf("No PowerZone for %d", pwr))
}

type PowerZone struct {
	Number   int
	MinWatts int
	MaxWatts int
}

func (s *PowerZone) Matches(power int) bool {
	return (s.MinWatts <= power) && (s.MaxWatts >= power)
}

func (s *PowerZone) AvgWatts() int {
	return (s.MaxWatts + s.MinWatts) / 2
}
