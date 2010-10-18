package common

import (
	"os"
	"regexp"
)

var (
	ADDRESS_REGEXP *regexp.Regexp
)


func init() {
	ADDRESS_REGEXP = regexp.MustCompile(generateAddressRegexp())
}

func generateAddressRegexp() string {
	digit := "([0-9])"
	optional_digit := "("+digit+"?)"
	decimal_byte := "("+digit+optional_digit+optional_digit+")"
	ip := "("+decimal_byte+"\\."+decimal_byte+"\\."+decimal_byte+"\\."+decimal_byte+")"
	port := "("+digit+optional_digit+optional_digit+optional_digit+optional_digit+")"
	return "^"+ip+"?:"+port+"$"
}

func IsRegularFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsRegular()
}

func IsValidAddress(addr string) bool {
	return ADDRESS_REGEXP.MatchString(addr)
}
