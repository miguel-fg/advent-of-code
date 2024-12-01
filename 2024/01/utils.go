package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ListAssembler(inputFilePath string) (list1 []int, list2 []int){

    f, err:= os.Open(inputFilePath)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        fields:= strings.Fields(line)

        f1, _ := strconv.Atoi(fields[0])
        f2, _ := strconv.Atoi(fields[1])

        list1 = append(list1, f1)
        list2 = append(list2, f2)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    sort.Ints(list1)
    sort.Ints(list2)

    return
}

func MapAssembler(list []int) map[int]int {
    hashMap := make(map[int]int)

    for i := 0; i < len(list); i++ {
        key := list[i] 
        hashMap[key]++
    }

    return hashMap
}

func TotalDistance(list1 []int, list2 []int) (total int){
    total = 0

    for i:= 0; i < len(list1); i++ {
        distance := absInt(list1[i] - list2[i])

        total += distance
    }

    return
}

func SimilarityScore(list []int, count map[int]int) (total int) {
    total = 0

    for i := 0; i < len(list); i++ {
        value := list[i]

        total += (value * count[value])
    }

    return
}

func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}


