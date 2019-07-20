package utils

import (
	"os/exec"
	"runtime"
)

func OpenLink(link string) {
	command := "open"
	if runtime.GOOS == "linux" {
		command = "x-www-browser"
	}
	if _, err := RunCommand(command, link); err != nil {
		Error.Println(err)
	}
}

func RunCommand(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	res, err := cmd.Output()
	if err != nil {
		Error.Println(err)
		return []byte{}, nil
	}
	return res, nil
}
