package main

import (
	"fmt"
	"os"
)

func main() {

	for key, val := range os.Args[1:] {
		fmt.Println(key, val)
	}
}
