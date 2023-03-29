package main

import (
	"fmt"
	"mad-aliens/simulation/cmd"
	"os"
)

func main() {
	cli := cmd.New()
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
