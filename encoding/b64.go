package encoding

import (
	"encoding/base64"
)

var B64 b64

type b64 struct{}

func (b64) Encode(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(buf, b)
	return buf
}

func (b64) Decode(b []byte) []byte {
	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, _ := base64.StdEncoding.Decode(dbuf, []byte(b))
	return dbuf[:n]
}

func (b64) EncodeStr(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (b64) DecodeStr(s string) string {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return string(decoded)
}
