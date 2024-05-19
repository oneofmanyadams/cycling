package cycling

import (
	"testing"
)

func TestAddTime(t *testing.T) {
	type testCase struct {
		vals     []int64
		want_len int
	}
	cases := []testCase{
		{vals: []int64{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []int64{}, want_len: 0},
		{vals: []int64{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddTime(v)
		}
		if len(tr.TimeStampEachSec) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.TimeStampEachSec), tc.want_len)
		}
		for vk, v := range tr.TimeStampEachSec {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %d, want %d", k, v, tc.vals[vk])
			}
		}
	}
}

func TestAddPower(t *testing.T) {
	type testCase struct {
		vals     []int
		want_len int
	}
	cases := []testCase{
		{vals: []int{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []int{}, want_len: 0},
		{vals: []int{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddPower(v)
		}
		if len(tr.PowerEachSec) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.PowerEachSec), tc.want_len)
		}
		for vk, v := range tr.PowerEachSec {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %d, want %d", k, v, tc.vals[vk])
			}
		}
	}
}

func TestAddHeartRate(t *testing.T) {
	type testCase struct {
		vals     []int
		want_len int
	}
	cases := []testCase{
		{vals: []int{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []int{}, want_len: 0},
		{vals: []int{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddHeartRate(v)
		}
		if len(tr.HeartRateEachSec) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.HeartRateEachSec), tc.want_len)
		}
		for vk, v := range tr.HeartRateEachSec {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %d, want %d", k, v, tc.vals[vk])
			}
		}
	}
}

func TestAddCadence(t *testing.T) {
	type testCase struct {
		vals     []int
		want_len int
	}
	cases := []testCase{
		{vals: []int{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []int{}, want_len: 0},
		{vals: []int{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddCadence(v)
		}
		if len(tr.CadenceEachSec) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.CadenceEachSec), tc.want_len)
		}
		for vk, v := range tr.CadenceEachSec {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %d, want %d", k, v, tc.vals[vk])
			}
		}
	}
}

func TestAddSpeed(t *testing.T) {
	type testCase struct {
		vals     []float64
		want_len int
	}
	cases := []testCase{
		{vals: []float64{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []float64{}, want_len: 0},
		{vals: []float64{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddSpeed(v)
		}
		if len(tr.SpeedEachSec) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.SpeedEachSec), tc.want_len)
		}
		for vk, v := range tr.SpeedEachSec {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %f, want %f", k, v, tc.vals[vk])
			}
		}
	}
}

func TestAddDistance(t *testing.T) {
	type testCase struct {
		vals     []float64
		want_len int
	}
	cases := []testCase{
		{vals: []float64{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []float64{}, want_len: 0},
		{vals: []float64{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddDistance(v)
		}
		if len(tr.DistanceCumulative) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.DistanceCumulative), tc.want_len)
		}
		for vk, v := range tr.DistanceCumulative {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %f, want %f", k, v, tc.vals[vk])
			}
		}
	}
}

func TestAddCalories(t *testing.T) {
	type testCase struct {
		vals     []int
		want_len int
	}
	cases := []testCase{
		{vals: []int{1, 2, 3, 4, 5}, want_len: 5},
		{vals: []int{}, want_len: 0},
		{vals: []int{5}, want_len: 1},
	}
	for k, tc := range cases {
		var tr Ride
		for _, v := range tc.vals {
			tr.AddCalories(v)
		}
		if len(tr.CaloriesCumulative) != tc.want_len {
			t.Fatalf("Test case %d incorrect # of vals added. got: %d, want: %d", k, len(tr.CaloriesCumulative), tc.want_len)
		}
		for vk, v := range tr.CaloriesCumulative {
			if v != tc.vals[vk] {
				t.Fatalf("Test case %d missmatch vals got: %d, want %d", k, v, tc.vals[vk])
			}
		}
	}
}
