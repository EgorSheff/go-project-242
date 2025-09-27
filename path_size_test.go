package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSizeFile(t *testing.T) {
	size, err := GetSize("testdata/sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(1064960), size)
}

func TestGetPathSizeDir(t *testing.T) {
	size, err := GetSize("testdata")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(3194880), size)
}

func TestFormat(t *testing.T) {
	testSet := map[string]int64{
		"3.0MB":  3194880,
		"30.1MB": 31581223,
		"2.9GB":  3158391223,
	}
	for expected, size := range testSet {
		require.Equal(t, expected, FormatSize(size, true))
	}
}
