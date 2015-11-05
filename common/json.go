/*
 Copyright (C) 2015 Enrico Bianchi (enrico.bianchi@gmail.com)
 Project       Memento
 Description   A backup system
 License       GPL version 2 (see GPL.txt for details)
*/

package common

import (
	"encoding/json"
	"net"
)

type JSONCommand struct {
	Name    string   `json:"name"`
	Cmd     string   `json:"cmd,omitempty"`
	Paths   []string `json:"paths,omitempty"`
	Element JSONFile `json:"element,omitempty"`
	ACL     bool     `json:"acl,omitempty"`
}

type JSONMessage struct {
	Context string      `json:"context"`
	Command JSONCommand `json:"command"`
}

type JSONResult struct {
	Result  string   `json:"result"`
	Message string   `json:"message,omitempty"`
	Data    JSONFile `json:"data,omitempty"`
}

type JSONFileAcl struct {
	Mode  string `json:"mode,omitempty"`
	User  string `json:"user,omitempty"`
	Group string `json:"group,omitempty"`
}

type JSONFile struct {
	Name       string        `json:"name"`
	Size       int64         `json:"size,omitempty"`
	Compressed bool          `json:"compressed,omitempty"`
	Hash       string        `json:"hash,omitempty"`
	Type       string        `json:"type"`
	Os         string        `json:"os"`
	Mode       string        `json:"mode,omitempty"`
	User       string        `json:"user,omitempty"`
	Group      string        `json:"group,omitempty"`
	Mtime      int64         `json:"mtime,omitempty"`
	Ctime      int64         `json:"ctime,omitempty"`
	Link       string        `json:"link,omitempty"`
	Acl        []JSONFileAcl `json:"acl,omitempty"`
}

func (j *JSONMessage) Send(conn net.Conn) error {
	jsonmessage, _ := json.Marshal(j)
	jsonmessage = append(jsonmessage, "\n"...)

	_, err := conn.Write(jsonmessage)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (j *JSONResult) Send(conn net.Conn) error {
	jsonmessage, _ := json.Marshal(j)
	jsonmessage = append(jsonmessage, "\n"...)

	_, err := conn.Write(jsonmessage)

	if err != nil {
		return err
	} else {
		return nil
	}
}
