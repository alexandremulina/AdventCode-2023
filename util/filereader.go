package util

import (
	"io"
	"io/ioutil"
	"strings"
)

func ReadFileContent(fileContent io.Reader) ([]string, error) {
	content, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return nil, err
	}

	contentString := string(content)

	var lines []string
	for _, line := range strings.Split(contentString, "\n") {
		lines = append(lines, strings.TrimSpace(line))
	}

	return lines, nil
}
