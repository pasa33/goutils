package encodingUtils

import jsoniter "github.com/json-iterator/go"

var Json thisJson
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type thisJson struct{}

func (thisJson) Unescape(s string) string {
	if len(s) == 0 {
		return ""
	}
	if string(s[0]) != `"` {
		s = `"` + s
	}
	if string(s[len(s)-1]) != `"` {
		s = s + `"`
	}
	t := `{"value":` + s + `}`
	jsonRes := map[string]string{}
	if json.Unmarshal([]byte(t), &jsonRes) != nil {
		return ""
	}
	return jsonRes["value"]
}
