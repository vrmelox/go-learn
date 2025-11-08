package main

import (
	"fmt"
	"os"
)


func verifySquare(mapi []string, row int, col int, size int) int {
	limiti := row + size
	limitj := col + size

	for i := row; i < limiti ; i++ {
		for j := col; j  < limitj ; j++ {
			if mapi[i][j] != '.' {
				return 0
			}
		}
	}
	return 1
}

// func checksquare(mapi []string, row int, col int, size int) int {
// 	cuti := mapi[row : row + size]
// 	for _, elem := range cuti {
// 		subs := elem[col : col + size]
// 		if strings.Contains(subs, "o") {
// 			return 0
// 		}
// 	}
// 	return 1
// }

func main () {
	file := os.Args[1]
	mapi, err := getContent(file)
	if err != nil {
		fmt.Println(err)
	}
	re := verifySquare(mapi, 5, 3, 4)
	fmt.Println(re)
}