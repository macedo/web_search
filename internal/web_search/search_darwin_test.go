//go:build darwin

package web_search

import (
	"testing"
)

func TestExecDarwin(t *testing.T) {
	e := NewGoogleEngine()
	s := New(e)

	command = mockCmd

	err := s.Exec()
	if err != nil {
		t.Errorf("expected to error, got %q instead", err)
	}
}
