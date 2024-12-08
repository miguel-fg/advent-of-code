package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func PartOne(inputFilePath string) int {
    equations := GetEquations(inputFilePath)
    result := CalculateEquations(equations)

    return result
}

func GetEquations(inputFilePath string) map[int][]int {

    f, err := os.Open(inputFilePath)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    equations := make(map[int][]int)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ": ")

        expected, _ := strconv.Atoi(parts[0])
        numbers := strings.Split(parts[1], " ")

        values := StrArrToIntArr(numbers)

        equations[expected] = values
    }

    if err := scanner.Err(); err != nil {
        log.Fatal (err)
    }

    return equations
}

func StrArrToIntArr(strArray []string) []int {
    var values []int

    for i := 0; i < len(strArray); i++ {
        val, _ := strconv.Atoi(strArray[i])

        values = append(values, val);
    }

    return values
}

func CalculateEquations(equations map[int][]int) int {
    result := 0

    for key, value := range equations {
        if ValidateEquation(key, value) {
            result += key
        }
    }

    return result
} 

func ValidateEquation(expected int, values []int) bool {
    if len(values) == 0 {
        return false
    }

    var dfs func(index int, current int) bool
    dfs = func(index int, current int) bool {
        if index == len(values) {
            return current == expected
        }

        return dfs(index + 1, current+values[index]) || dfs(index+1, current*values[index])
    }

    return dfs(1, values[0])
}
