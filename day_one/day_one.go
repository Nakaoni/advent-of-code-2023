package main

import (
	"github.com/Nakaoni/advent-of-code-2023/day_one/part_one"
	"github.com/Nakaoni/advent-of-code-2023/utils"
	"log"
)

func main() {
	log.SetPrefix("day_one: ")
	log.SetFlags(0)

	lines, err := utils.GetFileContent("./day_one/assets/input.txt")
	if err != nil {
		log.Panicln(err)
	}

	resultOne, errOne := part_one.GetResult(lines)
	//resultTwo, errTwo := part_two.GetResult(lines)

	utils.PrintResult("part_one", resultOne, errOne)
	//utils.PrintResult("part_two", resultTwo, errTwo)
}
