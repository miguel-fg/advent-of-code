package main

func PartTwo(fileInputPath string) int {
    matrix := WordSearchGenerator(fileInputPath)
    total := XMasFinder(matrix)
    
    return total
}

func XMasFinder(matrix [][]string) int {
    found := 0

    for i, row := range matrix {
	for j, col := range row {
	    if col == "A" {
		if D1Check(matrix, i, j) && D2Check(matrix, i, j){
		    found++
		}
	    }
	}
    }

    return found
}

func D1Check(matrix [][] string, row int, col int) bool {
    if(row + 1 >= len(matrix) || row - 1 < 0){
	return false
    }

    if(col + 1 >= len(matrix[row]) || col - 1 < 0){
	return false
    }

    valid := false 

    downLeft := matrix[row + 1][col - 1]
    upRight := matrix[row - 1][col + 1]

    if(downLeft == "M" && upRight == "S"){
	valid = true
    } else if (downLeft == "S" && upRight == "M") {
	valid = true
    }

    return valid
}

func D2Check(matrix [][] string, row int, col int) bool {
    if(row + 1 >= len(matrix) || row - 1 < 0){
	return false
    }

    if(col + 1 >= len(matrix[row]) || col - 1 < 0){
	return false
    }

    valid := false 

    upLeft := matrix[row - 1][col - 1]
    downRight := matrix[row + 1][col + 1]

    if(upLeft == "M" && downRight == "S") {
	valid = true
    } else if (upLeft == "S" && downRight == "M"){
	valid = true
    }

    return valid
}
