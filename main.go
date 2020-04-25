package main

import (
	cli "github.com/urfave/cli/v2"
	"os"
)

func main() {
	(&cli.App{}).Run(os.Args)
}
