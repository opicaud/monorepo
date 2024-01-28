package main

import (
	"context"
	"fmt"
	cli2 "github.com/opicaud/monorepo/shape-app/api/pacts/cli"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:  "create-default-shape",
		Usage: "make a rectangle ",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost",
				Usage: "host to run cli",
			},
			&cli.IntFlag{
				Name:  "port",
				Value: 50051,
				Usage: "host port",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			f := cli2.Client{}
			host := cmd.String("host")
			port := cmd.Int("port")
			area, err := f.CreateShape(fmt.Sprintf("%s:%d", host, port), CreateDefaultShapeRequest())
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			fmt.Printf("%d\n", area.GetCode())
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
