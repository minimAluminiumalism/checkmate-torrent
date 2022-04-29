package engine

import (
	"os"
	"strings"
	"github.com/minimAluminiumalism/simple-torrent/utils"
)

func moveFile(filePath string) error {
	fileStr := strings.Split(filePath, "/")
	fileName := fileStr[len(fileStr)-1]
	fileName = utils.Rename(fileName)
	err := os.Rename("downloads/" + filePath, "transfer/" + fileName)
	if err != nil {
		return err
	}
	return nil
}