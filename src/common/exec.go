/*
 Copyright (C) 2015 Enrico Bianchi (enrico.bianchi@gmail.com)
 Project       Memento
 Description   A backup system
 License       GPL version 2 (see GPL.txt for details)
*/

package common

import (
	"os/exec"
	"runtime"
	"bytes"
)


type cmdError struct {
	code    string
	message string
}

func (e *cmdError) Error() string {
	return e.message + " - " + e.code
}


// NOTE: the function is very dangerous, because is possible to create hangs on the system
func ExecuteCMD(command string) error {
	var cmd *exec.Cmd
	var err error

	stderr := &bytes.Buffer{}

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	cmd.Stderr = stderr

	if err = cmd.Run(); err != nil {
		// FIXME: due the specific software for now is not possible to manage exit with error from external program
		return &cmdError{err.Error(), string(stderr.Bytes())}
	}

	return nil
}
