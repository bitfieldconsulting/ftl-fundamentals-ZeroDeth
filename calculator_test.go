package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	name        string
	a, b        float64
	want        float64
  input       float64
  errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

// Introducing A slice of test Cases
func TestAddStruct(t *testing.T) {
	t.Parallel()
	testCases := []*testCase{
		{a: 2, b: 2, want: 4, name: "Two positive numbers that sum to a positive"},
		{a: -1, b: -1, want: -2, name: "Two negative numbers that sum ot a negative"},
		{a: -5, b: 0, want: -5, name: "Negative and positive number that sum Zero"},
	}

	// Looping over test cases
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

// Write a new test for one or more of your functions which generates random test inputs
// instead of using prepared cases (you can use the math/rand library for this).
func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		sum := a + b
		got := calculator.Add(a, b)
		if sum != got {
			t.Fatalf("Add(%f, %f) ; want %f, got %f", a, b, sum, got)
		}
	}
}

func TestSubtract(t *testing.T) {
  t.Parallel()
  testCases := []*testCase{
    {a: 2, b: 2, want: 0, name: "Two subtract by two are zero"},
    {a: 1, b: 2, want: -1, name: "One subtract by two are negative ones"},
    {a: 5, b: 0, want: 5, name: "Five subtract by zero are five"},
  }

  for _, tc := range testCases {
    got := calculator.Subtract(tc.a, tc.b)
    if tc.want != got {
      t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
    }
  }
}

func TestMultiply(t *testing.T) {
  t.Parallel()
  testCases := []*testCase{
    {a: 2, b: 2, want: 4, name: "Two multiply by two are four"},
    {a: 1, b: 2, want: 2, name: "One multiply by two are two"},
    {a: 5, b: 0, want: 0, name: "Five multiply by zero are zero"},
  }

  for _, tc := range testCases {
    got := calculator.Multiply(tc.a, tc.b)
    if tc.want != got {
      t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
    }
  }
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []*testCase{
		{a: 6, b: 0, want: 0, errExpected: true},
		{a: 4, b: 2, want: 5.5, errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, err)
		}
		if tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f) ; want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		a, b, want  float64
		errExpected bool
	}{
		{a: 9, want: 3, errExpected: false},
		{a: -4, want: 99, errExpected: true},
		{a: 14, want: 3.741657, errExpected: false},
		{a: -2, want: 99, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Sqrt(%f) unexpected error status: %v", tc.a, err)
		}
		if !tc.errExpected && !closeEnough(tc.want, got) {
			t.Errorf("Sqrt(%f) want %f, got %f", tc.a, tc.want, got)
		}
	}
}

func closeEnough(a, b float64) bool {
	var x = 0.01
	return math.Abs(a-b) < x
}
