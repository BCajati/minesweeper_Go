package main

import (
	"fmt"
	"strconv"
//	"strings"
)

type Square struct {
	isBomb bool
	state string
}

const(
	numRows = 5
	numCols = 5
)

var (

	board[numRows][numCols]Square

)

func createBoard(rows, cols int){
	var s1 Square
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++{
			s1 = Square{false,"X"}
			board[x][y]= s1
		}
	}

	// todo make this random
	board[1][1].isBomb = true
	board[3][3].isBomb = true
	board[4][2].isBomb = true
}

func showBoard(foundBomb bool){

	if foundBomb{
		fmt.Printf("You Lost! \n")
	}
	var x,y int
	for  x=0; x < numRows; x++ {
		for  y = 0; y < numCols; y++ {
			fmt.Printf(board[x][y].state)
		}
		fmt.Printf("\n")
	}
	
}

func searchRowforBombs(x,y, rowToSearch int) int {
	var numBombs = 0
	// if this square is not top row count bombs above
	var colStart, colEnd int 
	colStart = y
	colEnd = y
	if (y != 0){
		colStart = y-1
	}
	if (y != (numCols -1)) {
		colEnd = y + 1
	}
		var col int
		for col = colStart; col <= colEnd; col++ {
			if board[rowToSearch][col].isBomb {
				numBombs ++
			}
		}
	return numBombs
}

func findAdjacentBombs(x,y int) int {
	
	var numBombs = 0
	// search row above
	if (x != 0)	{
		numBombs = numBombs + searchRowforBombs(x,y, x-1)
	}

	// search row below
	if ( x != (numRows-1)) {
		numBombs = numBombs + searchRowforBombs(x,y, x+1)
	}

	// search same row
	// we can ignore this squares state since it will be zero,
	// otherwise we would have found a bomb here and lost the game
	numBombs = numBombs + searchRowforBombs(x,y, x)
	
	return numBombs
}

func pickSquare(x,y int) bool {

	fmt.Printf("Selected Square %d,%d \n", x,y)

	if board[x][y].isBomb{
		board[x][y].state = "*"
		return true
	}else{
		var numberOfAdjacentBombs int
		numberOfAdjacentBombs = findAdjacentBombs(x,y)
		board[x][y].state = strconv.Itoa(numberOfAdjacentBombs)
		return false
	}

}

func main() {
	
	fmt.Println("Welcome to Minesweeper")

	createBoard(numRows, numCols)
	
	showBoard(false)

	var foundBomb = pickSquare(1,0)
	showBoard(foundBomb)

	foundBomb = pickSquare(3,2)
	showBoard(foundBomb)

	foundBomb = pickSquare(3,4)
	showBoard(foundBomb)

	foundBomb = pickSquare(1,4)
	showBoard(foundBomb)


	foundBomb = pickSquare(0,2)
	showBoard(foundBomb)

	foundBomb = pickSquare(4,2)
	showBoard(foundBomb)

}