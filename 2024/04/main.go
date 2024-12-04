package main

import "fmt"

func main() {
	words := PartOne("input.txt")
    xmas := PartTwo("input.txt")
    fmt.Printf("WORDS: %d\n", words)
    fmt.Printf("X-MAS: %d\n", xmas)
}
