package web_search

import (
	"os/exec"
)

var command = exec.Command

func (s *Search) Exec() error {
	return command("open", s.engine.URL()).Run()
}
