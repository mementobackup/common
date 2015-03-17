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
)

// NOTE: the function is very dangerous, because is possible to create hangs on the system
func ExecuteCMD(command string) error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	if err := cmd.Run(); err != nil {
		// FIXME: due the specific software for now is not possible to manage exit with error from external program
	}

	return nil
}
