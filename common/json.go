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

type JSONAcl struct {
    Type  string `json:"type,omitempty"`
    Link  string `json:"link,omitempty"`
    Size  string `json:"size,omitempty"`
    Hash  string `json:"hash,omitempty"`
    Atime string `json:"atime,omitempty"`
    Mtime string `json:"mtime,omitempty"`
    Ctime string `json:"ctime,omitempty"`
    Mode  string `json:"mode,omitempty"`
    User  string `json:"user,omitempty"`
    Group string `json:"group,omitempty"`
}

type JSONCommand struct {
	Name      string `json:name`
	Directory string `json:directory`
	ACL       bool   `json:acl`
}

type JSONMessage struct {
	Context string      `json:context`
	Command JSONCommand `json:command`
}

type JSONResult struct {
    Result  string  `json:"result"`
    Message string  `json:"message"`
}

type JSONFile struct {
    Result  string  `json:result`
    Name    string  `json:"name,omitempty"`
    Os      string  `json:"os,omitempty"`
    Acl     JSONAcl `json:"acl,omitempty"`
}

func Sendresult(conn net.Conn, result JSONResult) {
    jsonmessage, _ := json.Marshal(result)

	conn.Write(jsonmessage)
}