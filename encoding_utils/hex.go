package encoding_utils

import (
	"encoding/hex"
	"fmt"
)

var Hex thishex

type thishex struct{}

func (thishex) Encode(b []byte) []byte {
	buf := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(buf, b)
	return buf
}

func (thishex) Decode(b []byte) []byte {
	dbuf := make([]byte, hex.DecodedLen(len(b)))
	n, _ := hex.Decode(dbuf, []byte(b))
	return dbuf[:n]
}

func (thishex) EncodeStr(s string) string {
	return hex.EncodeToString([]byte(s))
}

func (thishex) DecodeStr(s string) string {
	decoded, _ := hex.DecodeString(s)
	return string(decoded)
}

func (thishex) EncodeAny(s any, customPattern ...string) string {
	pattern := "%x"
	if len(customPattern) > 0 {
		pattern = customPattern[0]
	}
	return fmt.Sprintf(pattern, s)
}
