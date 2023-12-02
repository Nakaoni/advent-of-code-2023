package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetFileContent(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func PrintResult(prefix string, result string, err error) {
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("%v: %v\n", prefix, result)
}
