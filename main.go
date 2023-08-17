package main

import "fmt"

func drawBoard(grid [3][3]int) {      // Function taking a 3x3 grid, and drawing it on the console
    for i := 0; i < 3; i++ {          // Loop through each row in the grid
        for j := 0; j < 3; j++ {      // Loop through each square in the row
            square := grid[i][j]      // Assign the current value to square
            if square == 0 {          // If it is empty,
                fmt.Print(" ")        // Print a space
            } else if square == 1 {   // If it is player one,
                fmt.Print("X")        // Print "X"
            } else {                  // Otherwise,
                fmt.Print("O")        // Print "O"
            }                         // End if
            if j < 2 {                // After the first two columns,
                fmt.Print("|")        // Print "|"
            }                         // End if
        }                             // End for
        if i < 2 {                    // After the first two rows,
            fmt.Println("\n-+-+-")    // Print "-+-+-"
        }                             // End if
    }                                 // End for
}                                     // End func

func checkWin(grid [3][3]int) int {   // Function taking a 3x3 grid, and checking for a 3-in-a-row
    for x := 0; x < 3; x++ {          // Loop through each x from 0 to 2
        if grid[x][0] == grid[x][1] && grid[x][1] == grid[x][2] && grid[x][2] != 0 {
            return grid[x][0]         // Check for a completed row
        }                             // for either player 1 or 2
        if grid[0][x] == grid[1][x] && grid[1][x] == grid[2][x] && grid[2][x] != 0 {
            return grid[0][x]         // Check for a completed column
        }                             // for either player 1 or 2
    }                                 // End for
    if grid[0][0] == grid[1][1] && grid[1][1] == grid[2][2] && grid[2][2] != 0 {
        return grid[0][0]             // Check for a completed diagonal
    }                                 // for either player 1 or 2
    if grid[0][2] == grid[1][1] && grid[1][1] == grid[2][0] && grid[2][0] != 0 {
        return grid[0][2]             // Check for a completed anti-diagonal
    }                                 // for either player 1 or 2
    return 0                          // Otherwise, return 0
}                                     // End func

func checkDraw(grid [3][3]int) bool { // Function taking a 3x3 grid, and checking for a draw
    for a := 0; a < 3; a++ {          // Loop through each row in the grid
        for b := 0; b < 3; b++ {      // Loop through each square in the row
            if grid[a][b] == 0 {      // If the square is empty,
                return false          // Return false
            }                         // End if
        }                             // End for
    }                                 // End for
    return true                       // Otherwise, return true
}                                     // End func

func main() {                         // Main function taking no arguments
    board := [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
    curr_player := 1                  // Initialise variables: board is a 3x3 grid of ints,
    winner := 0                       // curr_player is an int, and winner is an int
    var input int                     // Initialise input variable as int
    for winner == 0 {                 // Loop while there is no winner
        fmt.Print("\033[H\033[2J")    // Clear the console
	    drawBoard(board)              // Draw the grid and print the current player
        fmt.Println("\n\nPlayer", curr_player, "\n")
        for true {                    // Forever,
            fmt.Scanln(&input)        // Scan for input
            if input > 0 && input < 10 {
                input --              // If it is a valid number and the squate is empty,
                if board[input / 3][input % 3] == 0 {
                    board[input / 3][input % 3] = curr_player
                    break             // Break from the loop
                }                     // End if
            }                         // End if
        }                             // End for
        curr_player = curr_player % 2 + 1
        winner = checkWin(board)      // Change the player and check for a winner
        if winner == 0 && checkDraw(board) {
            winner = -1               // If there is no winner but there is a draw,
        }                             // Set winner to -1
    }                                 // End for
    fmt.Print("\033[H\033[2J")        // Clear the console
    drawBoard(board)                  // Draw the grid
    if winner != -1 {                 // If it wasn't a draw, print the winner
        fmt.Println("\n\nPlayer", winner, "wins!")
    } else {                          // Otherwise,
        fmt.Println("\n\nDraw!")      // Print "Draw!"
    }                                 // End if
}                                     // End func

