package fmtwrapper

import (
	"os/exec"
	"bytes"
)

func Exec(name string, params ...string) string {
	cmd := exec.Command(name, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "ERROR"
	}
	return out.String()
}
