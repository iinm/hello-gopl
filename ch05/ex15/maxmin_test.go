package main

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 5},
		//{[]int{}, 5},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("max(%v)", test.values)
		got, err := max(test.values...)
		if err != nil || got != test.want {
			t.Errorf("%s = (%d, %v) want (%d, nil)", descr, got, err, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5, 4, 3, 2, 1}, 1},
		//{[]int{}, 5},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("min(%v)", test.values)
		got, err := min(test.values...)
		if err != nil || got != test.want {
			t.Errorf("%s = (%d, %v) want (%d, nil)", descr, got, err, test.want)
		}
	}
}

func TestMax2(t *testing.T) {
	tests := []struct {
		value  int
		values []int
		want   int
	}{
		{1, []int{2, 3, 4, 5}, 5},
		{5, []int{4, 3, 2, 1}, 5},
		//{5, []int{4, 3, 2, 1}, 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("max2(%d, %v)", test.value, test.values)
		got := max2(test.value, test.values...)
		if got != test.want {
			t.Errorf("%s = %d want %d", descr, got, test.want)
		}
	}
}

func TestMin2(t *testing.T) {
	tests := []struct {
		value  int
		values []int
		want   int
	}{
		{1, []int{2, 3, 4, 5}, 1},
		{5, []int{4, 3, 2, 1}, 1},
		//{5, []int{4, 3, 2, 1}, 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("min2(%d, %v)", test.value, test.values)
		got := min2(test.value, test.values...)
		if got != test.want {
			t.Errorf("%s = %d want %d", descr, got, test.want)
		}
	}
}
