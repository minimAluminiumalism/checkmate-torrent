package utils

import (
	"strings"
)


func Rename(oldName string) string {
	newName := oldName
	if strings.Contains(oldName, "hhd800.com@") {
		newName = strings.Replace(oldName, "hhd800.com@", "", -1)
	}
	if strings.Contains(oldName, "jav20s8.com@") {
		newName = strings.Replace(oldName, "jav20s8.com@", "", -1)
	}
	return newName
}
