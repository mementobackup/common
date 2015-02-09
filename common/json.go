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
    Name      string `json:name`
    Directory string `json:directory,omitempty`
    ACL       bool   `json:acl,omitempty`
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
    Type  string `json:"type"`
    Link  string `json:"link"`
    Size  string `json:"size"`
    Hash  string `json:"hash"`
    Atime string `json:"atime,omitempty"`
    Mtime string `json:"mtime,omitempty"`
    Ctime string `json:"ctime,omitempty"`
    Mode  string `json:"mode"`
    User  string `json:"user"`
    Group string `json:"group"`
}

type JSONFile struct {
    Result string        `json:"result"`
    Name   string        `json:"name"`
    Os     string        `json:"os"`
    Acl    []JSONFileAcl `json:"acl"`
}

func (j *JSONResult) Send(conn net.Conn) {
    jsonmessage, _ := json.Marshal(j)

    conn.Write(jsonmessage)
}

func (j *JSONFile) Send(conn net.Conn) {
    jsonmessage, _ := json.Marshal(j)

    conn.Write(jsonmessage)
}
