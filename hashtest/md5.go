package hashtest

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(fmt.Sprint("error:md5 os open error %v", err))
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", errors.New(fmt.Sprint("error:md5 io copy error %v", err))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

//get string md5
func StringMd5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}
