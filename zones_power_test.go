package cycling

import (
	"testing"
)

func TestNewPowerZones(t *testing.T) {
	type wantZone struct {
		num int
		min int
		max int
	}
	type testCase struct {
		ftp        int
		maxes      []float64
		want_zones []wantZone
	}
	cases := []testCase{
		{ftp: 250, maxes: []float64{0.55, 0.75, .90, 1.05, 1.20, 1.4, 3.0},
			want_zones: []wantZone{
				{num: 1, min: 0, max: 137},
				{num: 2, min: 138, max: 187},
				{num: 3, min: 188, max: 225},
				{num: 4, min: 226, max: 262},
				{num: 5, min: 263, max: 300},
				{num: 6, min: 301, max: 350},
				{num: 7, min: 351, max: 750},
			},
		},
		{ftp: 250, maxes: []float64{0.55, 0.75, .90, 1.10, 3.0},
			want_zones: []wantZone{
				{num: 1, min: 0, max: 137},
				{num: 2, min: 138, max: 187},
				{num: 3, min: 188, max: 225},
				{num: 4, min: 226, max: 275},
				{num: 5, min: 276, max: 750},
			},
		},
		{ftp: 250, maxes: []float64{0.75, 1.20, 3.0},
			want_zones: []wantZone{
				{num: 1, min: 0, max: 187},
				{num: 2, min: 188, max: 300},
				{num: 3, min: 301, max: 750},
			},
		},
	}
	for k, tc := range cases {
		got := NewPowerZones(tc.ftp, tc.maxes)
		if len(got) != len(tc.want_zones) {
			t.Fatalf("Test case %d wrong number of zones. got: %d, want: %d", k, len(got), len(tc.want_zones))
		}
		for tzk, tz := range got {
			if tz.MinWatts != tc.want_zones[tzk].min {
				t.Fatalf("Test case %d, zone %d non-matching MinWatts. got: %d, want %d", k, tzk, tz.MinWatts, tc.want_zones[tzk].min)
			}
			if tz.MaxWatts != tc.want_zones[tzk].max {
				t.Fatalf("Test case %d, zone %d non-matching MaxWatts. got: %d, want %d", k, tzk, tz.MaxWatts, tc.want_zones[tzk].max)
			}
		}
	}
}

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
