package utils

import (
	"fmt"
	"testing"
)

func TestRename(t *testing.T) {
	newName := Rename("dv-1224-uncensored.mp4")
	fmt.Println(newName)
}