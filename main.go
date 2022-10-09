package main

import (
	"fmt"
	"os"
	"github.com/nikola43/gohat/cli"
)

func main() {
	if err := cli.New().Run(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
