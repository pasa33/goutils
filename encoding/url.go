package encodingUtils

import (
	"net/url"
)

var Url thisUrl

type thisUrl struct{}

func (thisUrl) Encode(s string) string {
	return url.QueryEscape(s)
}

func (thisUrl) Decode(s string) string {
	decoded, _ := url.QueryUnescape(s)
	return decoded
}

func (thisUrl) FullEncode(s string) string {
	upperhex := "0123456789ABCDEF"
	llen := len(s)
	if llen == 0 {
		return s
	}

	var buf [64]byte
	var t []byte

	required := len(s) + 2*llen
	if required <= len(buf) {
		t = buf[:required]
	} else {
		t = make([]byte, required)
	}

	j := 0
	for i := 0; i < len(s); i++ {
		t[j] = '%'
		t[j+1] = upperhex[s[i]>>4]
		t[j+2] = upperhex[s[i]&15]
		j += 3
	}
	return string(t)
}
