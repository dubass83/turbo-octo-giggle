package common

import "testing"

func TestMax(t *testing.T) {
	expectedResult := 69
	actualResult := Max(2, 3, 45, 32, 69, 68, 33, 54, 11, 59)
	if expectedResult != actualResult {
		t.Errorf("function return unexpected result: %d", actualResult)
	}
}

func TestMaxOneValue(t *testing.T) {
	expectedResult := 69
	actualResult := Max(69)
	if expectedResult != actualResult {
		t.Errorf("function return unexpected result: %d", actualResult)
	}
}

func TestMin(t *testing.T) {
	expectedResult := -3
	actualResult := Min(2, -3, 45, 32, 69, 68, 33, 54, 11, 59)
	if expectedResult != actualResult {
		t.Errorf("function return unexpected result: %d", actualResult)
	}
}

func TestMinOneValue(t *testing.T) {
	expectedResult := 2
	actualResult := Min(2)
	if expectedResult != actualResult {
		t.Errorf("function return unexpected result: %d", actualResult)
	}
}
