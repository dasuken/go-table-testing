package main

import "testing"

var testTax tax

func TestMain(m *testing.M)  {
	setup()

	m.Run()
}

func setup() {
	testTax = newTax()
}

func TestTaxExcludeAmount(t *testing.T) {
	testCase := []struct {
		testName string
		input int
		want int
	}{
		{testName: "divided", input: 2160, want: 2000},
		{testName: "roundUp", input: 2005, want: 1856},
		{testName: "roundDown", input: 2000, want: 1852},
	}

	for _, tc := range testCase {
		t.Run(tc.testName, func(t *testing.T) {
			got := testTax.calculateTaxExcludeAmount(tc.input)
			if got != tc.want {
				t.Fatalf("want %v, but %v", tc.want, got)
			}
			t.Log("logger at ", t.Name())
		})
	}
}