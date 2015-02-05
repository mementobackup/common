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
    Name      string  `json:name`
    Directory string  `json:directory`
    ACL       boolean `json:acl`
}

type JSONMessage struct {
	Context string      `json:context`
	Command JSONCommand `json:command`
}

func Sendresult(conn net.Conn, result, message string) {
	type JSONResult struct {
		Result  string `json:"result"`
		Message string `json:"message"`
	}

	jsonmessage, _ := json.Marshal(&JSONResult{
		Result:  result,
		Message: message,
	})

	conn.Write(jsonmessage)
}
