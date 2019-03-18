package fmtwrapper

import (
	"os/exec"
	"bytes"
)

func Exec(name string, arg string) string {
	cmd := exec.Command(name, arg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	return out.String()
}
