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
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Value:   false,
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Value:   false,
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Value:   false,
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			size, err := code.GetPathSize(cmd.Args().First(), cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))
			if err != nil {
				return err
			}
			fmt.Println(size)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
