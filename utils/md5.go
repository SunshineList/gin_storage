package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 解密

func Md5Check(content []byte, encrypted string) bool {
	return strings.EqualFold(MD5V(content), encrypted)
}

// 加密

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
