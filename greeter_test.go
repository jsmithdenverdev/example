package main

import "testing"

func Test_Greeter(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		greeter  func(string) string
		expected string
	}{
		{
			testName: "english greeting",
			name:     "Jake",
			greeter:  englishGreeter,
			expected: "Howdy, Jake",
		},
	}

	for _, test := range tests {
		actual := test.greeter(test.name)
		if actual != test.expected {
			t.Errorf("%s, want: %s, got: %s", test.testName, test.expected, actual)
		}
	}
}
