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
		Usage: "print size of a file or directory;",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			size, err := code.GetSize(cmd.Args().First())
			if err != nil {
				return err
			}
			fmt.Printf("%dB\t%s\n", size, cmd.Args().First())
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
