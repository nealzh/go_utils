package file_utils

import (
	"io/ioutil"
	"log"
)

func LoadFileContent(filePath string) []byte {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}
