package main

import (
	"os"
	"fmt"
	)

func main() {
	path := os.Args
	err := os.Link(path[1], path[2])
	if err != nil {
		fmt.Println(err)
	}else {
		err := os.Remove(path[1])
		if err != nil {
			fmt.Println(err)
		}
	}
}