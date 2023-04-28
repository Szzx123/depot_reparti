package utils

import "strings"

func Findval(msg string, key string) string {
	if len(msg) < 4 {
		return ""
	}

	sep := msg[0:1]
	tab_allkeyvals := strings.Split(msg[1:], sep)

	for _, keyval := range tab_allkeyvals {
		if len(keyval) >= 4 {
			equ := keyval[0:1]
			tabkeyval := strings.Split(keyval[1:], equ)
			if tabkeyval[0] == key {
				return tabkeyval[1]
			}
		}
	}

	return ""
}
