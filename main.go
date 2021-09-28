package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	testFile, _ := filepath.Abs("../etc/test.html")
	data, err := os.ReadFile(testFile)
	check(err)

	s := string(data)
	fmt.Println(s)
}
