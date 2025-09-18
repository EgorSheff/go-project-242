package main

import (
	"code"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	recursive, human, all, help bool

	path string

	options = map[string]*bool{
		"--recursive": &recursive,
		"-r":          &recursive,

		"--human": &human,
		"-H":      &human,

		"--all": &all,
		"-a":    &all,

		"--help": &help,
		"-h":     &help,
	}
)

func main() {
	if err := parseArgs(); err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	size, err := code.CalcPathSize(path, recursive, human, all)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	fmt.Printf("%s %s\n", size, path)
}

// parseArgs - разбор параметров команды и установка соответствующих переменных
func parseArgs() error {
	for i, arg := range os.Args {
		// Пропускаем бинарник
		if i == 0 {
			continue
		}
		// Если параметр не начинается на -, то рассматриваем его как путь
		if !strings.HasPrefix(arg, "-") {
			path = arg
			continue
		}

		if opt, ok := options[arg]; ok {
			*opt = true
		} else {
			return fmt.Errorf("unsupported option: %s", arg)
		}
	}

	if path == "" {
		return errors.New("no path specified")
	}
	return nil
}
