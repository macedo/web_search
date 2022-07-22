//go:build linux

package web_search

import (
	"testing"
)

func TestExecLinux(t *testing.T) {
	e := NewGoogleEngine()
	s := New(e)

	mockIsWSLRelease = "f"
	isWSLRelease = func(r string) bool {
		return false
	}

	command = mockCmd

	err := s.Exec()
	if err != nil {
		t.Errorf("expected no error, got %q instead", err)
	}
}

func TestExecLinuxWSL(t *testing.T) {
	e := NewGoogleEngine()
	s := New(e)

	mockIsWSLRelease = "t"
	isWSLRelease = func(r string) bool {
		return true
	}

	command = mockCmd

	err := s.Exec()
	if err != nil {
		t.Errorf("expected no error, got %q instead", err)
	}
}
