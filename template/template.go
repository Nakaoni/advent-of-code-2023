package main

import (
	"github.com/Nakaoni/advent-of-code-2023/template/part_one"
	"github.com/Nakaoni/advent-of-code-2023/template/part_two"
	"github.com/Nakaoni/advent-of-code-2023/utils"
	"log"
)

func main() {
	log.SetPrefix("template: ")
	log.SetFlags(0)

	lines, err := utils.GetFileContent("./template/assets/input.txt")
	if err != nil {
		log.Panicln(err)
	}

	resultOne, errOne := part_one.GetResult(lines)
	resultTwo, errTwo := part_two.GetResult(lines)

	utils.PrintResult("part_one", resultOne, errOne)
	utils.PrintResult("part_ two", resultTwo, errTwo)
}
