package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSizeFile(t *testing.T) {
	size, err := GetSize("testdata/sample.txt", false, false)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(1064960), size)
}

func TestGetPathSizeDir(t *testing.T) {
	size, err := GetSize("testdata", false, false)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(1064960), size, "not hidden files")

	size, err = GetSize("testdata", true, false)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(3194880), size, "hidden files")
}

func TestGetPathSizeRecursive(t *testing.T) {
	size, err := GetSize("testdata", true, true)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int64(3201120), size)
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
