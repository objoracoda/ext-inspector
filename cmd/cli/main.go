package main

import (
	"fmt"
	"os"

	"github.com/objoracoda/ext-inspector/internal/cli"
)

func main() {
	args := os.Args[1:]

	if err := cli.Run(args); err != nil {
		fmt.Println("ошибка: ", err)
		return
	}
}
