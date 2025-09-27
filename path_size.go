package code

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var (
	units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, all, recursive)
	if err != nil {
		return "", err
	}
	fSize := FormatSize(size, human)

	return fSize, nil
}

func GetSize(path string, all, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var size int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") && !all {
			continue
		}

		if entry.IsDir() {
			if recursive {
				s, err := GetSize(filepath.Join(path, entry.Name()), all, recursive)
				if err != nil {
					return 0, err
				}
				size += s
			}
			continue
		}

		i, err := entry.Info()
		if err != nil {
			return 0, err
		}
		size += i.Size()
	}
	return size, nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	var baseIdx int
	for i := range units {
		base := math.Pow(2, 10*float64(i))
		if size < int64(base) {
			baseIdx = i - 1
			break
		}
	}
	if baseIdx == 0 {
		return fmt.Sprintf("%dB", size)
	}
	s := float64(size) / math.Pow(2, 10*float64(baseIdx))
	return fmt.Sprintf("%.1f%s", s, units[baseIdx])
}
