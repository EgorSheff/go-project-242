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
