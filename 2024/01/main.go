package main

import "fmt"

func main() {
    // PART 1
    l1, l2 := ListAssembler("input.txt")

    fmt.Printf("Total distance: %d\n", TotalDistance(l1, l2))

    // PART 2
    counts := MapAssembler(l2)

    fmt.Printf("Similarity score: %d", SimilarityScore(l1, counts))
}
