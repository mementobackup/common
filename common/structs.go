/*
 Copyright (C) 2015 Enrico Bianchi (enrico.bianchi@gmail.com)
 Project       Memento
 Description   A backup system
 License       GPL version 2 (see GPL.txt for details)
*/

package common

import "github.com/go-ini/ini"

type Section struct {
    Cfg     *ini.File
    Name    string
    Grace   string
    Dataset int
}
