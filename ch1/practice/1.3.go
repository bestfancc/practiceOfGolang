package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	for key, val := range os.Args[1:] {
		fmt.Println(key, val)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start1).Seconds())
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start2).Seconds())
}
