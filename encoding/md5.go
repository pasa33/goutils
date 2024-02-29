package encoding_utils

import (
	"crypto/md5"
	"encoding/hex"
)

var MD5 thismd5

type thismd5 struct{}

func (thismd5) Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
