package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileRead, err := os.ReadFile("./lines.txt")
	if err != nil {
		log.Fatal("can't read file for some reason ", fileRead)
	}
	trimFileRead := strings.TrimSpace(string(fileRead))
	lines := strings.Split(string(trimFileRead), "\n")
	result := 0
	for _, line := range lines {
		lineMulti := strings.Split(line, " * ")
		firstElement, err := strconv.Atoi(lineMulti[0])
		if err != nil {
			log.Fatal("can't conver for some reason", err)
		}
		secondElement, err := strconv.Atoi(lineMulti[1])
		resultLine := firstElement * secondElement
		result += resultLine
		fmt.Println(resultLine)
	}

	fmt.Println(result)
}
