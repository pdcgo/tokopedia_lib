package helper

import (
	"bufio"
	"os"
	"strings"
)

func FileLoadLineString(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func FileSaveLineString(fname string, lines []string) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	content := strings.Join(lines, "\n")
	_, err = file.Write([]byte(content))

	return err
}
