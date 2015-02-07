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
    Type  string `json:"type"`
    Link  string `json:"link"`
    Size  string `json:"size"`
    Hash  string `json:"hash"`
    Atime string `json:"atime"`
    Mtime string `json:"mtime"`
    Ctime string `json:"ctime"`
    Mode  string `json:"mode"`
    User  string `json:"user"`
    Group string `json:"group"`
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
    Name    string  `json:"name"`
    Os      string  `json:"os"`
    Acl     JSONAcl `json:"acl"`
}

func Sendresult(conn net.Conn, result &JSONResult) {
    jsonmessage, _ := json.Marshal(&JSONResult)

	conn.Write(jsonmessage)
}
