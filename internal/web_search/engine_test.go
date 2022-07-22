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

func TestDuckDuckGoEngineURL(t *testing.T) {
	e := NewDuckDuckGoEngine()
	e.SetQuery("golang")

	expectedURL := "https://duckduckgo.com?q=golang"

	if e.URL() != expectedURL {
		t.Errorf("expected URL %q, got %q instead", expectedURL, e.URL())
	}
}

func TestStackOverflowEngineURL(t *testing.T) {
	e := NewStackOverflowEngine()
	e.SetQuery("golang")

	expectedURL := "https://stackoverflow.com/search?q=golang"

	if e.URL() != expectedURL {
		t.Errorf("expected URL %q, got %q instead", expectedURL, e.URL())
	}
}
