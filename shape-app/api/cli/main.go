package main

import (
	"context"
	"fmt"
	"github.com/opicaud/monorepo/shape-app/api/pacts"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:  "create-default-shape",
		Usage: "make a rectangle ",
		Action: func(context.Context, *cli.Command) error {
			f := pacts.Foo{}
			area, err := f.GetRectangleAndSquareArea2(fmt.Sprintf("localhost:%d", 8080), CreateDefaultShapeRequest())
			fmt.Printf("%s, %s\n", area, err)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
