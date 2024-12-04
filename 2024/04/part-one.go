package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func PartOne(fileInputPath string) int{

    matrix := WordSearchGenerator(fileInputPath)
    total := WordSearchParser(matrix)

    return total
}

func WordSearchGenerator(fileInputPath string) [][]string {
    var wordSearch [][]string 
    
    f, err := os.Open(fileInputPath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        row := make([]string, len(line))

        for i, char := range line {
            row[i] = string(char)
        }

        wordSearch = append(wordSearch, row)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return wordSearch
}

func WordSearchParser(matrix [][]string) int {
    found := 0

    for i, row := range matrix {
        for j, col := range row {
            if col == "X" {
                if FwdCheck(matrix, i, j) {
                    found++
                }

                if BwdCheck(matrix, i, j) {
                    found++
                }

                if DownCheck(matrix, i, j) {
                    found++
                }

                if UpCheck(matrix, i, j) {
                    found++
                }

                if NECheck(matrix, i, j) {
                    found++
                }

                if SECheck(matrix, i, j) {
                    found++
                }

                if SWCheck(matrix, i, j) {
                    found++
                }

                if NWCheck(matrix, i, j) {
                    found++
                }
            }
        }
    }
    return found
}

func FwdCheck(matrix [][]string, row int, col int) bool {
    if(col + 3 >= len(matrix[row])){
        return false
    }

    return matrix[row][col + 1] == "M" && matrix[row][col + 2] == "A" && matrix[row][col + 3] == "S"
}

func BwdCheck(matrix [][] string, row int, col int) bool {
    if(col - 3 < 0){
        return false
    }

    return matrix[row][col - 1] == "M" && matrix[row][col - 2] == "A" && matrix[row][col - 3] == "S"
}

func DownCheck(matrix [][] string, row int, col int) bool {
    if(row + 3 >= len(matrix)) {
        return false
    }

    return matrix[row + 1][col] == "M" && matrix[row + 2][col] == "A" && matrix[row + 3][col] == "S"
}

func UpCheck(matrix [][] string, row int, col int) bool {
    if(row - 3 < 0) {
        return false
    }

    return matrix[row - 1][col] == "M" && matrix[row - 2][col] == "A" && matrix[row - 3][col] == "S"
}

func NECheck(matrix [][] string, row int, col int) bool {
    if(col + 3 >= len(matrix[row])){
        return false
    }

    if(row - 3 < 0) {
        return false
    }

    return matrix[row - 1][col + 1] == "M" && matrix[row - 2][col + 2] == "A" && matrix[row - 3][col + 3] == "S"
}

func SECheck(matrix [][] string, row int, col int) bool {
    if(col + 3 >= len(matrix[row])){
        return false
    }

    if(row + 3 >= len(matrix)){
        return false
    }

    return matrix[row + 1][col + 1] == "M" && matrix[row + 2][col + 2] == "A" && matrix[row + 3][col + 3] == "S"
}

func SWCheck(matrix [][] string, row int, col int) bool {
    if(col - 3 < 0) {
        return false
    }

    if(row + 3 >= len(matrix)){
        return false
    }

    return matrix[row + 1][col - 1] == "M" && matrix[row + 2][col - 2] == "A" && matrix[row + 3][col - 3] == "S"
}

func NWCheck(matrix [][] string, row int, col int) bool {
    if(col - 3 < 0){
        return false
    }

    if(row - 3 < 0){
        return false
    }

    return matrix[row - 1][col - 1] == "M" && matrix[row - 2][col - 2] == "A" && matrix[row - 3][col - 3] == "S"
}

func PrintMatrix(matrix [][]string) {
    for _, row := range matrix {
        for _, col := range row {
            fmt.Printf("%s ", col)
        }
        fmt.Println()
    }
}
