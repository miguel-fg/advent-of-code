package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
)

func PartOne(inputFilePath string) (safeReports int) {
    safeReports = 0

    f, err := os.Open(inputFilePath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        report := strings.Fields(line) 

        if isReportValid(report) {
            safeReports++
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return
}

func isReportValid(report []string) bool {

    i, j := 0, 1
    num1, _ := strconv.Atoi(report[i])
    num2, _ := strconv.Atoi(report[j])

    // check if incrementing or decrementing
    var incrementing bool
    if num1 < num2 {
        incrementing = true
    } else if num2 < num1 {
        incrementing = false
    } else {
        return false
    }

    for j < len(report) {
        // check difference between 1 and 3
        difference := absInt(num1 - num2)

        if difference == 0 || difference > 3 {
            return false
        }

        // not all levels incrementing
        if incrementing && (num2 < num1) {
            return false
        }

        // not all levels decrementing
        if !incrementing && (num1 < num2){
            return false
        }

        i++
        j++

        if j < len(report) {
            num1, _ = strconv.Atoi(report[i])
            num2, _ = strconv.Atoi(report[j])
        }
    }

    return true
}

func absInt(x int) int {
    if(x < 0) {
        return -x
    }

    return x
}
