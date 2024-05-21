package cycling

type Zones []Zone

type Zone struct {
	Level       int
	Name        string
	Description string
	MinDuration int
	MaxDuration int
}
