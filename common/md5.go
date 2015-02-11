/*
 Copyright (C) 2015 Enrico Bianchi (enrico.bianchi@gmail.com)
 Project       Memento
 Description   A backup system
 License       GPL version 2 (see GPL.txt for details)
*/

package common

import (
	"crypto/md5"
	"io"
	"math"
	"os"
)

const filechunk = 8192

func Md5(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	info, _ := file.Stat()
	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize - int64(i * filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	return hash.Sum(nil)
}
