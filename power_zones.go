package cycling

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
