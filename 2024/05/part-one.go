package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
	"strconv"
)

func PartOne(rulesFilePath string, updatesFilePath string) int {

    rules := GetRules(rulesFilePath)
    total := UpdateParser(updatesFilePath, rules)
    
    return total
}

func GetRules(inputFilePath string) map[string][]string {
    rulesMap := make(map[string][]string)

    f, err := os.Open(inputFilePath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        ruleset := strings.Split(line, "|")

        r1 := ruleset[0]
        r2 := ruleset[1]

        rulesMap[r2] = append(rulesMap[r2], r1)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return rulesMap
}

func UpdateParser(inputFilePath string, rules map[string][]string) int {
    total := 0

    f, err := os.Open(inputFilePath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        update := strings.Split(line, ",")

        total += GetUpdateValue(update, rules)
    }

    return total
}

func GetUpdateValue(update []string, rules map[string][]string) int {
    var visited []string
    visited = append(visited, update[0])

    for i := 1; i < len(update); i++ {
        page := update[i]
        for _, value := range visited {
            if !slices.Contains(rules[page], value) {
                return 0
            }
        }
        visited = append(visited, page)
    }  

    middle := (len(update) - 1)/2
    value, _ := strconv.Atoi(update[middle])

    return value
}
