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
	Name      string   `json:name`
	Value     string   `json:value,omitempty`
	Directory []string `json:directory,omitempty`
	Filename  string   `json:filename,omitempty`
	ACL       bool     `json:acl,omitempty`
}

type JSONMessage struct {
	Context string      `json:context`
	Command JSONCommand `json:command`
}

type JSONResult struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type JSONFileAcl struct {
	Mode  string `json:"mode,omitempty"`
	User  string `json:"user,omitempty"`
	Group string `json:"group,omitempty"`
}

type JSONFile struct {
	Result string        `json:"result"`
	Name   string        `json:"name"`
	Size   int64         `json:"size,omitempty"`
	Hash   string        `json:"hash,omitempty"`
	Type   string        `json:"type"`
	Os     string        `json:"os"`
	Mode   string        `json:"mode,omitempty"`
	User   string        `json:"user,omitempty"`
	Group  string        `json:"group,omitempty"`
	Mtime  int64         `json:"mtime,omitempty"`
	Acl    []JSONFileAcl `json:"acl,omitempty"`
}

func (j *JSONResult) Send(conn net.Conn) {
	jsonmessage, _ := json.Marshal(j)
	jsonmessage = append(jsonmessage, "\n"...)

	conn.Write(jsonmessage)
}

func (j *JSONFile) Send(conn net.Conn) {
	jsonmessage, _ := json.Marshal(j)
	jsonmessage = append(jsonmessage, "\n"...)

	conn.Write(jsonmessage)
}
