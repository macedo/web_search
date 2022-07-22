package web_search

import (
	"bytes"
	"os/exec"
	"regexp"
)

var command = exec.Command

var isWSLRelease = func(release string) bool {
	matched, _ := regexp.Match(`.*microsoft.*`, []byte(release))
	return matched
}

func (s *Search) Exec() error {
	cmd := command("xdg-open", s.engine.URL())

	release, err := kernelRelease()
	if err != nil {
		return err
	}

	if isWSLRelease(release) {
		cmd = command("cmd.exe", "/c", "start", s.engine.URL())
	}

	return cmd.Run()
}

func kernelRelease() (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("uname", "-r")
	cmd.Stdout = &out

	err := cmd.Run()
	return out.String(), err
}
