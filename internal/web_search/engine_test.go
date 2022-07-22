package web_search

import "testing"

func TestGoogleEngineURL(t *testing.T) {
	e := NewGoogleEngine()
	e.SetQuery("golang")

	expectedURL := "https://google.com/search?q=golang"

	if e.URL() != expectedURL {
		t.Errorf("expected URL %q, got %q instead", expectedURL, e.URL())
	}
}
