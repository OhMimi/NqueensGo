package main

import "fmt"

func main() {
	queen(8)
}

func queen(size int) {
	boards := make([]int, size)
	put(boards, 0)
}

func put(boards []int, col int) {
	size := len(boards)
	if col == size {
		fmt.Println(boards)
		for _, q := range boards {
			for k := 0; k < size; k++ {
				if k == q {
					fmt.Print("Q")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
		return
	}

	for pos := 0; pos < size; pos++ {
		boards[col] = pos
		if safe(boards, col) {
			put(boards, col+1)
		}
	}
}

func safe(boards []int, col int) bool {
	for c := 0; c < col; c++ {
		if isAttack(boards, c, col) {
			return false
		}
	}
	return true
}

func isAttack(boards []int, c int, col int) bool {
	if boards[c] == boards[col] {
		return true
	}
	if boards[col]-boards[c] == c-col {
		return true
	}
	if boards[col]-boards[c] == col-c {
		return true
	}
	return false
}
