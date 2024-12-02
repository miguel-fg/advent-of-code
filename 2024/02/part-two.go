package main

import (
	"bufio"
	"log"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func PartTwo(inputFilePath string) (safeReports int) {
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
        fmt.Printf("Processing report: %v\n", report)

        if IsReportValid(report) {
            fmt.Println("Report is valid without changes.")
            safeReports++
            continue
        }

        anomaly := AnomalyFinder(report)

        if anomaly != -1 {
            fmt.Printf("Anomaly found at index: %d\n", anomaly)
            reportOpt1 := removeFromReport(report, anomaly)
            reportOpt2 := removeFromReport(report, anomaly - 1)

            fmt.Printf("Option 1: %v\n", reportOpt1)
            fmt.Printf("Option 2: %v\n", reportOpt2)

            if IsReportValid(reportOpt1) || IsReportValid(reportOpt2) {
                fmt.Println("Report is valid after fixing anomaly.")
                safeReports++
            } else {
                fmt.Println("Neither option fixes the report")
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return
}

func AnomalyFinder(report []string) int {
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
        return j
    }

    for j < len(report) {
        // check difference between 1 and 3
        difference := AbsInt(num1 - num2)

        if difference == 0 || difference > 3 {
            return j
        }

        // not all levels incrementing
        if incrementing && (num2 < num1) {
            return j
        }

        // not all levels decrementing
        if !incrementing && (num1 < num2){
            return j
        }

        i++
        j++

        if j < len(report) {
            num1, _ = strconv.Atoi(report[i])
            num2, _ = strconv.Atoi(report[j])
        }
    }

    return -1
}

func removeFromReport(report []string, index int) []string {
    if index < 0 || index >= len(report) {
        return report
    }

    newReport := make([]string, 0, len(report) -1)
    newReport = append(newReport, report[:index]...)
    newReport = append(newReport, report[index + 1:]...)

    return newReport
}
