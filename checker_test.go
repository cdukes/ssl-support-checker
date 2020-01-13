package checker

import "testing"

func TestMaybeUpgradeURL(t *testing.T) {
	var (
		s        string
		expected string
	)

	s = MaybeUpgradeURL("http://example.com")
	expected = "http://example.com"

	if s != expected {
		t.Errorf("URL is incorrect, got: %s, expected: %s.", s, expected)
	}

	s = MaybeUpgradeURL("http://google.com")
	expected = "https://google.com"

	if s != expected {
		t.Errorf("URL is incorrect, got: %s, expected: %s.", s, expected)
	}
}
