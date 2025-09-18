package code

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var (
	humanSizes = [7]string{"", "KB", "MB", "GB", "TB", "PB", "EB"}
)

func CalcPathSize(path string, recursive, human, all bool) (string, error) {
	rootInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	rootAbsPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// Если файл
	if !rootInfo.IsDir() {
		if !all && rootInfo.Name()[0] == '.' {
			return "0", nil
		}

		if human {
			return getHumanSize(rootInfo.Size()), nil
		}
		return fmt.Sprint(rootInfo.Size()), nil
	}

	var size int64
	if err := filepath.WalkDir(rootAbsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relDir := filepath.Dir(strings.TrimPrefix(path, rootAbsPath))

		// Пропускаем скрытые файлы если не стоит нужная настройка
		if !all && strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		// Пропускаем вхождения которые не находятся в корне папки если не указан флаг -r
		if !recursive && relDir != "/" {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		if !d.IsDir() {
			size += info.Size()
			return nil
		}

		return nil
	}); err != nil {
		return "", err
	}

	if human {
		return getHumanSize(size), nil
	}

	return fmt.Sprint(size), nil
}

func getHumanSize(bytes int64) (human string) {
	for i, u := range humanSizes {
		base := math.Pow(2, float64(i)*10)

		if bytes < int64(base) {
			return
		}

		human = fmt.Sprintf("%.1f %s", float64(bytes)/base, u)
	}
	return
}
