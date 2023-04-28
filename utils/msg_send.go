package utils

import (
	"fmt"
)

var fieldsep = "/"
var keyvalsep = "="

func Msg_format(key string, val string) string {
	return fieldsep + keyvalsep + key + keyvalsep + val
}

func Msg_send(msg string) {
	fmt.Print(msg + "\n")
}
