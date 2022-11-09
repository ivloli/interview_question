package rangeList

import (
	"testing"
)

func TestAddRange(t *testing.T) {
	a := RangeList{
		interSlice: []int{},
	}
	a.AddRange(1, 5)
	result := a.ToString()
	respect := "[1,5)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	a.AddRange(0, 6)
	result = a.ToString()
	respect = "[0,6)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	a.AddRange(9, 15)
	result = a.ToString()
	respect = "[0,6) [9,15)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	a.AddRange(10, 11)
	result = a.ToString()
	respect = "[0,6) [9,15)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	a.AddRange(5, 7)
	result = a.ToString()
	respect = "[0,7) [9,15)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	a.AddRange(10, 18)
	result = a.ToString()
	respect = "[0,7) [9,18)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}
}

func TestRemoveRange(t *testing.T) {
	a := []int{0, 5, 8, 12, 9, 20}
	r := RangeList{a}

	r.RemoveRange(1, 4)
	result := r.ToString()
	respect := "[0,1) [4,5) [8,12) [9,20)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	r.RemoveRange(6, 7)
	result = r.ToString()
	respect = "[0,1) [4,5) [8,12) [9,20)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	r.RemoveRange(5, 9)
	result = r.ToString()
	respect = "[0,1) [4,5) [9,12) [9,20)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}

	r.RemoveRange(-1, 21)
	result = r.ToString()
	respect = "[)"
	if result != respect {
		t.Errorf("failed: result: %s, respect: %s \n", result, respect)
	}
}
