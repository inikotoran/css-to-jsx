package main

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			css := c.Args().Get(0)
			jsx, err := transform(css)
			if err != nil {
				fmt.Println("somethings wrong:", err)
			} else {
				fmt.Println(jsx)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
