package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a file")
	} else {
		fileName := os.Args[1]
		if fileName == "" {
			fmt.Println("Please provide a file")
		} else {
			tokens := lexFile(fileName)
			parsedTokens := parseTokens(tokens)
			symbolTable := createSymbolTable(parsedTokens)
			fmt.Println(parsedTokens)
			fmt.Println(symbolTable)
		}
	}
}

func lexFile(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	rawLines := string(data)
	parsedLines := strings.Split(rawLines, "\n")
	return parsedLines
}

func parseTokens(tokens []string) [][]string {
	parsedTokens := make([][]string, len(tokens))
	for i := 0; i < len(tokens); i++ {
		pieces := strings.Split(tokens[i], " ")
		parsedTokens[i] = pieces
	}
	return parsedTokens
}

func createSymbolTable(lines [][]string) map[int][]string {
	symbolTable := make(map[int][]string)
	for i := 0; i < len(lines); i++ {
		if strings.ContainsAny(lines[i][0], "0123456789") {
			lineNum, err := strconv.Atoi(lines[i][0])
			if err != nil {
				panic("Incorrect Line Num Format at line" + strconv.Itoa(i))
			}
			symbolTable[lineNum] = lines[i][1:]
		}
	}
	return symbolTable
}
