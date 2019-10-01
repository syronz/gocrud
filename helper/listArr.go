package helper

import (
	"fmt"
)

func ListArr(arr []string, msg string) {
	fmt.Println(msg)
	for i, v := range arr {
		fmt.Printf("%v. %v \n", i, v)
	}
}
