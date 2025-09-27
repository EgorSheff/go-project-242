package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "human",
				Value: false,
				Usage: "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:  "all",
				Value: false,
				Usage: "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			size, err := code.GetSize(cmd.Args().First(), cmd.Bool("all"))
			if err != nil {
				return err
			}
			fSize := code.FormatSize(size, cmd.Bool("human"))

			fmt.Printf("%s\t%s\n", fSize, cmd.Args().First())
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
