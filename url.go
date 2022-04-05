package utils

import (
	"strings"
)

//QueryString for string
func QueryString(m map[string]string) string {

	var rs []string
	for i, v := range m {
		rs = append(rs, i+"="+v)
	}

	return strings.Join(rs, "&")
}
