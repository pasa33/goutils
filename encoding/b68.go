package encodingUtils

import "github.com/shengdoushi/base58"

var B58 b58

type b58 struct{}

func (b58) Encode(s string, customAlphabet ...string) string {
	var alph *base58.Alphabet
	if len(customAlphabet) == 0 {
		alph = base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	} else {
		alph = base58.NewAlphabet(customAlphabet[0])
	}
	if customAlphabet == nil {
		return ""
	}
	return base58.Encode([]byte(s), alph)
}

func (b58) Decode(s string, customAlphabet ...string) string {
	var alph *base58.Alphabet
	if len(customAlphabet) == 0 {
		alph = base58.NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	} else {
		alph = base58.NewAlphabet(customAlphabet[0])
	}
	if customAlphabet == nil {
		return ""
	}
	txt, _ := base58.Decode(s, alph)
	return string(txt)
}
