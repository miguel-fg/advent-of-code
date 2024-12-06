package main

import (
    "bufio"
    "os"
    "log"
    "fmt"
    "strings"
)

func PartOne(inputFilePath string) int {

    guardMap := GenerateMap(inputFilePath)
    guardPosition := FindGuard(guardMap)
    StartPatrol(guardMap, guardPosition[0], guardPosition[1])

    return CountVisited(guardMap)
}

func GenerateMap(inputFilePath string) [][]string {
    f, err := os.Open(inputFilePath)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    var guardMap [][]string

    for scanner.Scan() {
        line := scanner.Text()
        chars := strings.Split(line, "")

        guardMap = append(guardMap, chars)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return guardMap
}

func FindGuard(guardMap [][]string) [2]int {
    var coords [2]int
    guard := "^"

    for i := 0; i < len(guardMap); i++ {
        for j := 0; j < len(guardMap[i]); j++ {
            if guardMap[i][j] == guard {
                coords[0] = i
                coords[1] = j

                return coords
            }
        }
    }

    return coords
}

func StartPatrol(guardMap [][]string, startRow int, startCol int) {
    const (
        Up string = "up"
        Down string = "down"
        Left string = "left"
        Right string = "right"
    )

    direction := Up
    currRow := startRow 
    currCol := startCol

    MarkVisited(guardMap, currRow, currCol) // starting position

    for {
        nextRow, nextCol := currRow, currCol

        switch direction {
        case Up:
            nextRow = currRow - 1
        case Down:
            nextRow = currRow + 1
        case Left:
            nextCol = currCol - 1
        case Right:
            nextCol = currCol + 1
        }

        if !Next(guardMap, nextRow, nextCol) {
            break
        }

        if guardMap[nextRow][nextCol] == "#" {
            switch direction {
            case Up:
                direction = Right
            case Down:
                direction = Left
            case Left:
                direction = Up
            case Right:
                direction = Down
            }
            continue
        }

        currRow, currCol = nextRow, nextCol
        MarkVisited(guardMap, currRow, currCol)
    }
}

func MarkVisited(guardMap [][]string, row int, col int){
    if row < 0 || row > len(guardMap) {
        return
    }

    if col < 0 || col > len(guardMap[row]) {
        return
    }

    if(guardMap[row][col] != "X") {
        guardMap[row][col] = "X"
    }
}

func Next(guardMap [][]string, row int, col int) bool {
    return row >= 0 && row < len(guardMap) && col >= 0 && col < len(guardMap[row])
}

func CountVisited(guardMap [][]string) int {
    total := 0

    for i := 0; i < len(guardMap); i++ {
        for j := 0; j < len(guardMap[i]); j++ {
            if guardMap[i][j] == "X" {
                total++
            }
        }
    }

    return total
}

func PrintMap(guardMap [][]string) {
    for i := 0; i < len(guardMap); i++ {
        for j := 0; j < len(guardMap[i]); j++ {
            fmt.Printf("%s ", guardMap[i][j])
        }
        fmt.Println()
    }
}
