package cycling

import (
	"fmt"
)

type PowerZone struct {
	Number      int
	Name        string
	Description string
	MinWatts    int
	MaxWatts    int
}

type PowerZones struct {
	Zones []PowerZone
}

func (s *PowerZones) DisplayZones() {
	for _, zone := range s.Zones {
		fmt.Println("Zone #%1", zone.Number)
	}
}
