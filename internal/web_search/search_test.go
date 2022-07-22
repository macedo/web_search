package web_search

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

var mockIsWSLRelease = "f"

func TestExecCommandHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	cmdName := ""

	switch runtime.GOOS {
	case "linux":
		cmdName = "xdg-open"

		v, err := strconv.ParseBool(os.Getenv("GO_WANT_WSL_KERNEL_RELEASE"))
		if err != nil {
			v = false
		}

		if v {
			cmdName = "cmd.exe"
		}
	case "darwin":
		cmdName = "open"
	}

	if strings.Contains(os.Args[2], cmdName) {
		os.Exit(0)
	}

	os.Exit(1)
}

func mockCmd(exe string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecCommandHelper"}
	cs = append(cs, exe)
	cs = append(cs, args...)

	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{
		"GO_WANT_HELPER_PROCESS=1",
		"GO_WANT_WSL_KERNEL_RELEASE=" + mockIsWSLRelease,
	}

	return cmd
}
