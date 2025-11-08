package main

import "fmt"


func verifySquare(mapi []string, row int, col int, size int) int {
	i := row
	j := col
	for ; i < i + size ; i++ {
		for ; j  < j + size ; j++ {
			if mapi[i][j] != '.' {
				return 0
			}
		}
	}
	return 1
}

