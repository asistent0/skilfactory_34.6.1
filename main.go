package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filenameOutput := "./output.txt"

	_ = os.Remove(filenameOutput)
	outputFile, err := os.OpenFile(filenameOutput, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(outputFile)

	filenameInput := "./input.txt"

	content, err := os.ReadFile(filenameInput)
	if err != nil {
		panic(err)
	}

	contentLines := strings.Split(string(content), "\n")

	re, _ := regexp.Compile(`([0-9]+)([+-/*])([0-9])+=+\?+`)

	for i := 0; i < len(contentLines); i++ {
		if contentLines[i] != "" && re.FindString(contentLines[i]) != "" {
			fmt.Println(contentLines[i])

			subMatch := re.FindAllStringSubmatch(contentLines[i], -1)
			fmt.Println(subMatch[0][1], subMatch[0][2], subMatch[0][3])

			number1, err := strconv.ParseFloat(subMatch[0][1], 64)
			number2, err := strconv.ParseFloat(subMatch[0][3], 64)

			_, err = writer.WriteString(
				subMatch[0][1] +
					subMatch[0][2] +
					subMatch[0][3] +
					"=" +
					floatToString(calculator(number1, subMatch[0][2], number2)) +
					"\n")
			if err != nil {
				panic(err)
			}
		}
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}
}

func floatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 0, 64)
}

func calculator(number1 float64, operator string, number2 float64) float64 {
	var result float64
	if operator == "-" {
		result = number1 - number2
	}
	if operator == "+" {
		result = number1 + number2
	}
	if operator == "*" {
		result = number1 * number2
	}
	if operator == "/" && number2 != 0 {
		result = number1 / number2
	}

	return result
}
