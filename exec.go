package fmtwrapper

import (
	"os/exec"
	"bytes"
)

func Exec(name string) string {
	cmd := exec.Command(name)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "ERROR"
	}
	return out.String()
}
