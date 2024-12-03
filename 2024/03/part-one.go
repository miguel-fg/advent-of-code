package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
    "strconv"
)

func PartOne(inputFilePath string) int {
    operations := MemoryCleaner(inputFilePath)

    result := RunProgram(operations)

    return result
}

func MemoryCleaner(inputFilePath string) []string {
    content, err := os.ReadFile(inputFilePath)

    if err != nil {
        log.Fatal(err)
    }

    r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

    matches := r.FindAllString(string(content), -1)

    return matches
}

func RunProgram(operations []string) (total int) {
    total = 0

    for _, op := range operations {
        total += Mul(op)
    }

    return
}

func Mul(instruction string) int {
    total := 1
    r := regexp.MustCompile(`[0-9]{1,3}`)
    operands := r.FindAllString(instruction, -1)

    if len(operands) == 2{
        num1, _ := strconv.Atoi(operands[0])
        num2, _ := strconv.Atoi(operands[1])

        total = num1 * num2

    } else {
        fmt.Printf("Error at instruction: %s\nOperands: 5%v\n", instruction, operands)
        log.Fatal("Failed to extract operands from instruction")
    }

    return total
}
