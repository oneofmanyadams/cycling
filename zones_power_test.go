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
func TestMatches_Power(t *testing.T) {
	type testCase struct {
		tst_pwr   int
		z_min_pwr int
		z_max_pwr int
		want      bool
	}
	cases := []testCase{
		{tst_pwr: 250, z_min_pwr: 240, z_max_pwr: 260, want: true},
		{tst_pwr: 250, z_min_pwr: 250, z_max_pwr: 260, want: true},
		{tst_pwr: 250, z_min_pwr: 240, z_max_pwr: 250, want: true},
		{tst_pwr: 250, z_min_pwr: 260, z_max_pwr: 270, want: false},
		{tst_pwr: 250, z_min_pwr: 230, z_max_pwr: 240, want: false},
		{tst_pwr: 0, z_min_pwr: 0, z_max_pwr: 240, want: true}}
	for k, tc := range cases {
		t_zone := PowerZone{MinWatts: tc.z_min_pwr, MaxWatts: tc.z_max_pwr}
		got := t_zone.Matches(tc.tst_pwr)
		if tc.want != got {
			t.Fatalf("Test case %d got: %t, want: %t", k, got, tc.want)
		}
	}
}

func TestAvgWatts(t *testing.T) {
	type testCase struct {
		z_min_pwr int
		z_max_pwr int
		want      int
	}
	cases := []testCase{
		{z_min_pwr: 240, z_max_pwr: 260, want: 250},
		{z_min_pwr: 235, z_max_pwr: 260, want: 247},
		{z_min_pwr: 220, z_max_pwr: 260, want: 240},
		{z_min_pwr: 0, z_max_pwr: 160, want: 80},
		{z_min_pwr: 270, z_max_pwr: 2060, want: 1165}}
	for k, tc := range cases {
		t_zone := PowerZone{MinWatts: tc.z_min_pwr, MaxWatts: tc.z_max_pwr}
		got := t_zone.AvgWatts()
		if tc.want != got {
			t.Fatalf("Test case %d got: %d, want: %d", k, got, tc.want)
		}
	}
}
