/*
 Copyright (C) 2015 Enrico Bianchi (enrico.bianchi@gmail.com)
 Project       Memento
 Description   A backup system
 License       GPL version 2 (see GPL.txt for details)
*/

package common

import (
	"io"
	"net"
	"os"
	"strings"
)

const BUFFER = 1024

func Sendfile(filename string, connection net.Conn) error {
	var err error

	file, err := os.Open(strings.TrimSpace(filename)) // For read access.
	if err != nil {
		connection.Write([]byte("-1"))
		return err
	}
	defer file.Close()

	_, err = io.Copy(connection, file)
	if err != nil {
		return err
	}

	return nil
}
