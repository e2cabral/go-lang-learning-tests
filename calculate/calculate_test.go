package calculate_test

import (
	"test-bible/calculate"
	"testing"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("Should return true for 371", func(t *testing.T) {
		testCase := TestCase{
			value:    371,
			expected: true,
		}

		testCase.actual = calculate.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("Should return true for 153", func(t *testing.T) {
		testCase := TestCase{
			value:    153,
			expected: true,
		}

		testCase.actual = calculate.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    250,
		expected: false,
	}

	testCase.actual = calculate.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
