package main

import "testing"

func TestMedianFinder(t *testing.T) {
	var mf MedianFinder
	if _, ok := mf.Median(); ok {
		t.Error("empty stream should have no median")
	}
	steps := []struct {
		add  int
		want float64
	}{
		{5, 5},
		{15, 10},
		{1, 5},
		{3, 4},
		{8, 5},
		{7, 6},
		{9, 7},
		{10, 7.5},
	}
	for _, s := range steps {
		mf.Add(s.add)
		got, ok := mf.Median()
		if !ok || got != s.want {
			t.Errorf("after Add(%d): median = %v, want %v", s.add, got, s.want)
		}
	}
}

func TestDescendingInserts(t *testing.T) {
	var mf MedianFinder
	for v := 10; v >= 1; v-- {
		mf.Add(v)
	}
	if got, _ := mf.Median(); got != 5.5 {
		t.Errorf("median of 1..10 = %v, want 5.5", got)
	}
}
