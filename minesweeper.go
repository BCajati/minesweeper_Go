package main

import (
	"fmt"
//	"strings"
)

type Square struct {
	X int
	Y int
	isBomb bool
	state string
}

var (
	s1 = Square{0, 0, false, "X"}  
	s2 = Square{0,1,false,"X"}  
	s3 = Square{0,2,false,"X"} 
	s4 = Square{1,0,false,"X"} 
	s5 = Square{1,1,false,"X"} 
	s6 = Square{1,2,true,"X"} 
	s7 = Square{2,0,false,"X"} 
	s8 = Square{2,1,false,"X"} 
	s9 = Square{2,2,false,"X"} 
//	board = [][]string{
//		[]string{"X", "X", "X"},
//		[]string{"X", "X", "X"},
//		[]string{"X", "X", "X"},
//	}

	board = [][]Square{
		[]Square{s1, s2, s3},
		[]Square{s4, s5, s6},
		[]Square{s7, s8, s9},
	}
)

func showBoard(foundBomb bool){
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}

	if foundBomb{
		fmt.Printf("You Lost!")
	}
		fmt.Printf("%s %s %s \n", board[0][0].state, board[0][1].state, board[0][2].state)
		fmt.Printf("%s %s %s \n", board[1][0].state, board[1][1].state, board[1][2].state)
		fmt.Printf("%s %s %s \n", s7.state, s8.state, s9.state)
	
}

func pickSquare(x,y int) bool {

	//var s Square := board[x][y]
	// s := board[x][y]

	if board[x][y].isBomb{
		board[x][y].state = "*"
		return true
	}else{
		// TODO need to compute values
		board[x][y].state = "1" 
		return false
	}
}

func main() {
	
	fmt.Println("Welcome to Minesweeper")
	
	showBoard(false)

	fmt.Println("Select Square 1,0")
	var foundBomb = pickSquare(1,0)

	showBoard(foundBomb)



}