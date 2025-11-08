package main

import "testing"

func TestGetH1fromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "empty string",
			html:     "",
			expected: "",
		},
		{
			name:     "simple h1",
			html:     "<h1>Welcome to My Site</h1>",
			expected: "Welcome to My Site",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1fromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: '%v', actual: '%v'", i, tc.name, tc.expected, actual)
			}
		})
	}
}