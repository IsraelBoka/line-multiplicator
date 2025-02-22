package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calcul(lines []string) (result int) {
	for _, line := range lines {
		lineMulti := strings.Split(line, " * ")
		firstElement, err := strconv.Atoi(lineMulti[0])
		if err != nil {
			log.Fatal("can't convert the first element for some reason", err)
		}
		secondElement, err := strconv.Atoi(lineMulti[1])

		if err != nil {
			log.Fatal("can't convert the second element for some reason", err)
		}
		resultLine := firstElement * secondElement
		result += resultLine
	}
	return
}

func main() {
	fileRead, err := os.ReadFile("./lines.txt")
	if err != nil {
		log.Fatal("can't read file for some reason ", fileRead)
	}
	trimFileRead := strings.TrimSpace(string(fileRead))
	lines := strings.Split(string(trimFileRead), "\n")
	result := Calcul(lines)
	fmt.Println("resultat final ", result)
}
