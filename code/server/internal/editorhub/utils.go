package editorhub

import (
	"bufio"
	"os"
)

func readFileContent(fileName string) ([][]byte, error) {
	var content [][]byte

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		content = append(content, lineCopy)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}
