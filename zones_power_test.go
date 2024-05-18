package cycling

import (
	"testing"
)

func TestMatchZone(t *testing.T) {
	zones := PowerZones{}
	zones = append(zones, PowerZone{Number: 1, MinWatts: 0, MaxWatts: 100})
	zones = append(zones, PowerZone{Number: 2, MinWatts: 101, MaxWatts: 200})
	zones = append(zones, PowerZone{Number: 3, MinWatts: 201, MaxWatts: 300})
	zones = append(zones, PowerZone{Number: 4, MinWatts: 301, MaxWatts: 400})
	zones = append(zones, PowerZone{Number: 5, MinWatts: 401, MaxWatts: 500})
	type testCase struct {
		pwr      int
		want     int
		want_err bool
	}
	cases := []testCase{
		{95, 1, false},
		{195, 2, false},
		{245, 3, false},
		{325, 4, false},
		{475, 5, false},
		{595, 0, true},
		{0, 1, false}}
	for _, tc := range cases {
		got, err := zones.MatchZone(tc.pwr)
		if got.Number != tc.want {
			t.Fatalf("got: %d, want: %d", got.Number, tc.want)
		}
		if err != nil && tc.want_err == false {
			t.Fatalf("Unexpected error returned: %s", err.Error())
		}
		if err == nil && tc.want_err == true {
			t.Fatalf("Expected error for pwr: %d, got %d, want: %d", tc.pwr, tc.want, got.Number)
		}
	}
}
