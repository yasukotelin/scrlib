package main

import (
	"fmt"
	"os"
	"time"

	"github.com/yasukotelin/scrlib"
)

func main() {
	fmt.Println("aaaaa")
	fmt.Println("aaaaa")
	fmt.Println("aaaaa")
	fmt.Println("aaaaa")
	fmt.Println("aaaaa")
	fmt.Println("clear this message after 3 seconds.")

	time.Sleep(3 * time.Second)

	if err := scrlib.Clear(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
