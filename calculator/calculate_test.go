package calculator_test

import (
	"test-bible/calculator"
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

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("Should return true for 153", func(t *testing.T) {
		testCase := TestCase{
			value:    153,
			expected: true,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("Should return false to 250", func(t *testing.T) {
		testCase := TestCase{
			value:    250,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("Should return false to 649", func(t *testing.T) {
		testCase := TestCase{
			value:    649,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
