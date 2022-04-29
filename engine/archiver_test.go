package engine

import "testing"

func TestMoveFile(t *testing.T) {
	err := moveFile("test/hhd800.com@STARS-542.mp4")
	if err != nil {
		t.Error(err)
	}
}