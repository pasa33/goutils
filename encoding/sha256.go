package encodingUtils

import (
	"crypto/sha256"
	"encoding/hex"
)

var Sha256 thisSha256

type thisSha256 struct{}

func (thisSha256) Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
