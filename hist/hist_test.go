package hist

import (
	"testing"
)

func TestMinAndMax(t *testing.T) {
	h := NewHistogram(10, 1)
	h.Add(5)

	percentiles := h.Percentiles(0.0, 1.0)
	if percentiles[0] != 5 {
		t.Errorf("Percentile(0.0) == %d (should be %d)\n", percentiles[0], 5)
	}
	if percentiles[1] != 5 {
		t.Errorf("Percentile(1.0) == %d (should be %d)\n", percentiles[1], 5)
	}
}

func TestDistinctValues(t *testing.T) {
	h := NewHistogram(100, 1)
	for i := 1; i <= 100; i++ {
		h.Add(i)
	}

	ps := h.Percentiles(0.0, 0.01, 0.02, 0.95, 0.99, 1.0)
	expected := []int{1, 1, 2, 95, 99, 100}
	for i, p := range ps {
		if p != expected[i] {
			t.Errorf("Actual(%d) != Expected(%d)", p, expected[i])
		}
	}
}

func TestOverlappingValues(t *testing.T) {
	h := NewHistogram(10, 1)
	for i := 1; i < 10; i++ {
		h.Add(i)
		h.Add(i)
	}

	ps := h.Percentiles(0.06, 0.07, 0.08, 0.09, 0.1, 0.17)
	expected := []int{1, 1, 1, 1, 1, 2}
	for i, p := range ps {
		if p != expected[i] {
			t.Errorf("Actual(%d) != Expected(%d)", p, expected[i])
		}
	}
}

func TestErrorCountAndRate(t *testing.T) {
	h := NewHistogram(10, 1)
	for i := 1; i < 10; i++ {
		h.AddError(1)
	}

	if h.ErrorPercent() != 100.0 {
		t.Errorf("Actual(%f) != Expected(%f)", h.ErrorPercent(), 100.0)
	}
}
