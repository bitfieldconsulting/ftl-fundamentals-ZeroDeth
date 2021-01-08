package calculator_test

import (
	"calculator"
	"testing"
)

//
type testCase struct {
	a, b        float64
	want        float64
	name        string
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

func TestAdd1and1(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Add(1, 1)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

// Introducing A slice of test Cases
func TestAddStruct(t *testing.T) {
	t.Parallel()
	testCases := []*testCase{
		{a: 2, b: 2, want: 4, name: "Two postive numbers that sum to a positive"},
		{a: -1, b: -1, want: -2, name: "Two negative numbers that sum ot a negative"},
		{a: -5, b: 0, want: 0, name: "Negative and positive number that sum Zero"},
	}

	// Loopig over test cases
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, tc.name, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Multiply(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

// Todo: InComplete TestDivide function
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
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)

			// A shorter way to write this logic is:
			/*
			    if tc.errExpected != (err != nil) {
			    t.Fatalf(...)
			   }
			*/
		}
		if tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f) ; want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}
