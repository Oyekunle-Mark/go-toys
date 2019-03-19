package main

import (
	"fmt"
	"flag"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to be read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Contents of file: ", string(data))
}