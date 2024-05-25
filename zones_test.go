package cycling

import "testing"

func TestLen(t *testing.T) {
	type testCase struct {
		tzns Zones
		want int
	}
	cases := []testCase{
		{tzns: Zones{}, want: 0},
		{tzns: Zones{{Level: 1}}, want: 1},
		{tzns: Zones{{Level: 1}, {Level: 2}}, want: 2},
		{tzns: Zones{{Level: 1}, {Level: 2}, {Level: 3}}, want: 3},
	}
	for k, tc := range cases {
		got := tc.tzns.Len()
		if got != tc.want {
			t.Fatalf("Test %d got %d want %d", k, got, tc.want)
		}
	}
}

func TestLess(t *testing.T) {
	type testCase struct {
		tzns  Zones
		small int
		big   int
		want  bool
	}
	cases := []testCase{
		{tzns: Zones{
			{Level: 1}, {Level: 2}},
			small: 0, big: 1, want: true},
		{tzns: Zones{
			{Level: 1}, {Level: 2}},
			small: 1, big: 0, want: false},
		{tzns: Zones{
			{Level: 1}, {Level: 2}},
			small: 1, big: 3, want: false},
	}
	for k, tc := range cases {
		got := tc.tzns.Less(tc.small, tc.big)
		if got != tc.want {
			t.Fatalf("Test %d got: %t, want: %t", k, got, tc.want)
		}
	}
}

func TestSwap(t *testing.T) {
	type testCase struct {
		tzns Zones
		key1 int
		key2 int
	}
	cases := []testCase{
		{tzns: Zones{
			{Level: 1, Name: "One"},
			{Level: 2, Name: "Two"}},
			key1: 0, key2: 1},
		{tzns: Zones{
			{Level: 1, Name: "One"},
			{Level: 2, Name: "Two"},
			{Level: 3, Name: "Three"}},
			key1: 0, key2: 2},
	}
	for k, tc := range cases {
		name1 := tc.tzns[tc.key1].Name
		name2 := tc.tzns[tc.key2].Name
		tc.tzns.Swap(tc.key1, tc.key2)
		if tc.tzns[tc.key1].Name != name2 {
			t.Fatalf("Test %d got: %s, want: %s", k, tc.tzns[tc.key1].Name, name2)
		}
		if tc.tzns[tc.key2].Name != name1 {
			t.Fatalf("Test %d got: %s, want: %s", k, tc.tzns[tc.key2].Name, name1)
		}
	}

}

func TestMaxHeartRates(t *testing.T) {
	type testCase struct {
		want []float64
	}
	cases := []testCase{
		{want: []float64{0.5, 1.0, 1.5, 2.0}},
	}
	for k, tc := range cases {
		var zns Zones
		for _, w := range tc.want {
			zns = append(zns, Zone{MaxHeartRate: w})
		}
		max_hrs := zns.MaxHeartRates()
		if len(max_hrs) != len(tc.want) {
			t.Fatalf("Test %d missmatch lens got: %d, want: %d", k, len(max_hrs), len(tc.want))
		}
		for key, got := range max_hrs {
			want_val := tc.want[key]
			if got != want_val {
				t.Fatalf("Test %d got: %f, want: %f", k, got, want_val)
			}
		}
	}
}

func TestMaxPowers(t *testing.T) {
	type testCase struct {
		want []float64
	}
	cases := []testCase{
		{want: []float64{0.5, 1.0, 1.5, 2.0}},
	}
	for k, tc := range cases {
		var zns Zones
		for _, w := range tc.want {
			zns = append(zns, Zone{MaxPower: w})
		}
		max_hrs := zns.MaxPowers()
		if len(max_hrs) != len(tc.want) {
			t.Fatalf("Test %d missmatch lens got: %d, want: %d", k, len(max_hrs), len(tc.want))
		}
		for key, got := range max_hrs {
			want_val := tc.want[key]
			if got != want_val {
				t.Fatalf("Test %d got: %f, want: %f", k, got, want_val)
			}
		}
	}
}
