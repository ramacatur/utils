package utils

import (
	"regexp"
	"strconv"
	"strings"
)

//GetNumFromString get 0-9 from string
func GetNumFromString(s string) string {
	re := regexp.MustCompile("[^0-9]+")

	return re.ReplaceAllString(s, "")
}

//IsNumeric func
func IsNumeric(s string) bool {
	_, err := strconv.Atoi(strings.Join(regexp.MustCompile(`[0-9]`).FindAllString(s, -1), ""))

	return err == nil
}

//IsObjectID func
func IsObjectID(s string) bool {
	match, _ := regexp.MatchString("^[0-9a-fA-F]{24}$", s)

	return match
}
