package debug_utils

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func PrintAsJson(v any, title ...string) {
	if len(title) > 0 {
		fmt.Printf("%s [%s] %s\n", strings.Repeat("=", 20), title[0], strings.Repeat("=", 20))
	}

	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
