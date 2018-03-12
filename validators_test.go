package main

import "testing"

func TestValidateURL(t *testing.T) {
	var urlTests = []struct {
		input    string
		expected string
	}{
		{"foobar", ""},
		{"http://example.com/?foobar=1#id", "http://example.com/?foobar=1#id"},
		{"example.com/test", "http://example.com/test"},
		{"https://example.com/test", "https://example.com/test"},
		{"hts://example.com/test", ""},
		{"  example.com/test  ", "http://example.com/test"},
		{"---", ""},
	}

	for _, pt := range urlTests {
		actual, _ := validateURL(pt.input)
		if actual != pt.expected {
			t.Fatalf("validateURL(%+v): expected %s, actual %s", pt.input, pt.expected, actual)
		}
	}

}
