package main

import (
	"fmt"
	"github.com/Nakaoni/advent-of-code-2023/template/part_one"
	"github.com/Nakaoni/advent-of-code-2023/template/part_two"
	"log"
	"os"
)

func main() {
	log.SetPrefix("template: ")
	log.SetFlags(0)

	_, err := GetFileContent("./template/assets/input.txt")
	if err != nil {
		log.Panicln(err)
	}

	resultOne, errOne := part_one.GetResult("Part 1")
	resultTwo, errTwo := part_two.GetResult("Part 2")

	printResult("part_one", resultOne, errOne)
	printResult("part_ two", resultTwo, errTwo)
}

func GetFileContent(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(file))

	return file, nil
}

func printResult(prefix string, result string, err error) {
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("%v: %v\n", prefix, result)
}
