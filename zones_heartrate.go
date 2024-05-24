package cycling

import (
	"errors"
	"fmt"
	"sort"
)

type HeartRateZones []HeartZone

func NewHeartRateZones(fthr int, maxhrs []float64) HeartRateZones {
	sort.Float64s(maxhrs)
	var hrzs HeartRateZones
	for len(hrzs) < len(maxhrs) {
		var new_zone HeartZone
		new_zone.Number = len(hrzs) + 1
		if new_zone.Number == 1 {
			new_zone.MinBPM = 0
		} else {
			new_zone.MinBPM = hrzs[len(hrzs)-1].MaxBPM + 1
		}
		new_zone.MaxBPM = int(float64(fthr) * maxhrs[len(hrzs)])
		hrzs = append(hrzs, new_zone)
	}
	return hrzs
}

func (s *HeartRateZones) MatchZone(hr int) (HeartZone, error) {
	for _, zone := range *s {
		if zone.Matches(hr) {
			return zone, nil
		}
	}
	return HeartZone{}, errors.New(fmt.Sprintf("No HeartZone for %d", hr))
}

type HeartZone struct {
	Number int
	MinBPM int
	MaxBPM int
}

func (s *HeartZone) Matches(heart_rate int) bool {
	return (s.MinBPM <= heart_rate) && (s.MaxBPM >= heart_rate)
}

func (s *HeartZone) AvgBPM() int {
	return (s.MaxBPM + s.MinBPM) / 2
}
