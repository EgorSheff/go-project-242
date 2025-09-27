package code

import "os"

func GetSize(path string) (int64, error) {
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
		if entry.IsDir() {
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
