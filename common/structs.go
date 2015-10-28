/*
 Copyright (C) 2015 Enrico Bianchi (enrico.bianchi@gmail.com)
 Project       Memento
 Description   A backup system
 License       GPL version 2 (see GPL.txt for details)
*/

package common

import "fmt"

type Section struct {
	Name       string
	Grace      string
	Dataset    int
	Compressed bool
}

type OperationErr struct {
	Position  int
	Operation string
	Message   string
}

func (e *OperationErr) Error() string {
	return fmt.Sprintf(" %s: %d - %s", e.Operation, e.Position, e.Message)
}
