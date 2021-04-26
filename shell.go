package main

import (
	"os"
	"os/exec"
)

var (
	c      *exec.Cmd
	err    error
	result []byte
)

func execute(cmd string) string {
	c = exec.Command("bash", "-c", cmd)
	result, err = c.Output()
	if err != nil {
		os.Exit(1)
	}
	return string(result)
}