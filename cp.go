package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main() {
	path := os.Args
	file, err := os.Open(path[1])
	if err != nil {
		fmt.Println("Can't open file")
		return
	}
	data := make([]byte, 0)
	for {
		buf := make([]byte, 4096)
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading from the file.")
			return
		}
		data = append(data, buf[:n]...)
		if n < 4096 {
			break
		}
	}
	// file2, err := os.Open(path[2])
	// if err != nil {
	// 	fmt.Println("Can't copy to second file")
	// 	return
	// }
	ioutil.WriteFile(path[2], data, 0644)
	fmt.Println(string(data))
}