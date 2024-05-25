package cycling

import (
	"testing"
)

func TestNewHeartRateZones(t *testing.T) {
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
		{ftp: 160, maxes: []float64{0.68, 0.83, .94, 1.05, 1.15, 1.2, 1.3},
			want_zones: []wantZone{
				{num: 1, min: 0, max: 108},
				{num: 2, min: 109, max: 132},
				{num: 3, min: 133, max: 150},
				{num: 4, min: 151, max: 168},
				{num: 5, min: 169, max: 184},
				{num: 6, min: 185, max: 192},
				{num: 7, min: 193, max: 208},
			},
		},
		{ftp: 160, maxes: []float64{0.68, 0.83, 1.00, 1.15, 1.3},
			want_zones: []wantZone{
				{num: 1, min: 0, max: 108},
				{num: 2, min: 109, max: 132},
				{num: 3, min: 133, max: 160},
				{num: 4, min: 161, max: 184},
				{num: 5, min: 185, max: 208},
			},
		},
		{ftp: 160, maxes: []float64{0.70, 1.00, 1.3},
			want_zones: []wantZone{
				{num: 1, min: 0, max: 112},
				{num: 2, min: 113, max: 160},
				{num: 3, min: 161, max: 208},
			},
		},
	}
	for k, tc := range cases {
		got := NewHeartRateZones(tc.ftp, tc.maxes)
		if len(got) != len(tc.want_zones) {
			t.Fatalf("Test case %d wrong number of zones. got: %d, want: %d", k, len(got), len(tc.want_zones))
		}
		for tzk, tz := range got {
			if tz.MinBPM != tc.want_zones[tzk].min {
				t.Fatalf("Test case %d, zone %d non-matching MinBPM. got: %d, want %d", k, tzk, tz.MinBPM, tc.want_zones[tzk].min)
			}
			if tz.MaxBPM != tc.want_zones[tzk].max {
				t.Fatalf("Test case %d, zone %d non-matching MaxBPM. got: %d, want %d", k, tzk, tz.MaxBPM, tc.want_zones[tzk].max)
			}
		}
	}
}

func TestMatchZone_HeartRate(t *testing.T) {
	zones := HeartRateZones{}
	zones = append(zones, HeartZone{Number: 1, MinBPM: 0, MaxBPM: 100})
	zones = append(zones, HeartZone{Number: 2, MinBPM: 101, MaxBPM: 200})
	zones = append(zones, HeartZone{Number: 3, MinBPM: 201, MaxBPM: 300})
	zones = append(zones, HeartZone{Number: 4, MinBPM: 301, MaxBPM: 400})
	zones = append(zones, HeartZone{Number: 5, MinBPM: 401, MaxBPM: 500})
	type testCase struct {
		hr       int
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
		got, err := zones.MatchZone(tc.hr)
		if got.Number != tc.want {
			t.Fatalf("got: %d, want: %d", got.Number, tc.want)
		}
		if err != nil && tc.want_err == false {
			t.Fatalf("Unexpected error returned: %s", err.Error())
		}
		if err == nil && tc.want_err == true {
			t.Fatalf("Expected error for hr: %d, got %d, want: %d", tc.hr, tc.want, got.Number)
		}
	}
}
func TestMatches_HeartRate(t *testing.T) {
	type testCase struct {
		tst_hr   int
		z_min_hr int
		z_max_hr int
		want     bool
	}
	cases := []testCase{
		{tst_hr: 160, z_min_hr: 150, z_max_hr: 170, want: true},
		{tst_hr: 160, z_min_hr: 160, z_max_hr: 170, want: true},
		{tst_hr: 160, z_min_hr: 150, z_max_hr: 160, want: true},
		{tst_hr: 160, z_min_hr: 170, z_max_hr: 180, want: false},
		{tst_hr: 160, z_min_hr: 140, z_max_hr: 150, want: false},
		{tst_hr: 0, z_min_hr: 0, z_max_hr: 170, want: true}}
	for k, tc := range cases {
		t_zone := HeartZone{MinBPM: tc.z_min_hr, MaxBPM: tc.z_max_hr}
		got := t_zone.Matches(tc.tst_hr)
		if tc.want != got {
			t.Fatalf("Test case %d got: %t, want: %t", k, got, tc.want)
		}
	}
}

func TestAvgBPM(t *testing.T) {
	type testCase struct {
		z_min_hr int
		z_max_hr int
		want     int
	}
	cases := []testCase{
		{z_min_hr: 140, z_max_hr: 160, want: 150},
		{z_min_hr: 135, z_max_hr: 160, want: 147},
		{z_min_hr: 120, z_max_hr: 160, want: 140},
		{z_min_hr: 0, z_max_hr: 160, want: 80},
		{z_min_hr: 100, z_max_hr: 200, want: 150}}
	for k, tc := range cases {
		t_zone := HeartZone{MinBPM: tc.z_min_hr, MaxBPM: tc.z_max_hr}
		got := t_zone.AvgBPM()
		if tc.want != got {
			t.Fatalf("Test case %d got: %d, want: %d", k, got, tc.want)
		}
	}
}
