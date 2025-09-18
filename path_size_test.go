package code

import (
	"fmt"
	"os"
	"testing"
)

func TestHumanSize(t *testing.T) {
	testSet := map[int64]string{
		1_300: "1.3 KB",
	}

	for b, base := range testSet {
		res := getHumanSize(b)
		if res != base {
			t.Fatalf("failed pair: bytes %d != %s", b, res)
		}
	}
}

func TestSingleFileCalc(t *testing.T) {
	// Single file
	if err := createZeroFile("test.txt", 1e10); err != nil {
		t.Fatal(err)
	}
	if err := createZeroFile(".test.txt", 1e10); err != nil {
		t.Fatal(err)
	}

	s, err := CalcPathSize("test.txt", false, false, false)
	if err != nil {
		t.Fatal(err)
	}
	sH, err := CalcPathSize("test.txt", false, true, false)
	if err != nil {
		t.Fatal(err)
	}
	sHidden, err := CalcPathSize(".test.txt", false, false, false)
	if err != nil {
		t.Fatal(err)
	}
	sHiddenOn, err := CalcPathSize(".test.txt", false, false, true)
	if err != nil {
		t.Fatal(err)
	}

	if s != fmt.Sprintf("%.0f", 1e10) {
		t.Fatalf("Size: sizes do not match: %.0f %s", 1e10, s)
	}
	if sH != "9.3 GB" {
		t.Fatalf("Size Human: sizes do not match: %s %s", "9,3GB", s)
	}
	if sHidden != "0" {
		t.Fatalf("Size noHidden: sizes do not match: 0 %s", s)
	}
	if sHiddenOn != fmt.Sprintf("%.0f", 1e10) {
		t.Fatalf("Size Hidden: sizes do not match: %.0f %s", 1e10, s)
	}

	errHandler(os.Remove("test.txt"), t)
	errHandler(os.Remove(".test.txt"), t)
}

func TestDirCalc(t *testing.T) {
	errHandler(os.Mkdir("testDir", 0777), t)
	errHandler(createZeroFile("testDir/test.txt", 1e5), t)
	errHandler(createZeroFile("testDir/test2.txt", 1e5), t)
	errHandler(createZeroFile("testDir/.test.txt", 1e5), t)

	errHandler(os.Mkdir("testDir/int", 0777), t)
	errHandler(createZeroFile("testDir/int/test3.txt", 1e4), t)

	recSize, err := CalcPathSize("testDir", true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	size, err := CalcPathSize("testDir/", false, true, false)
	if err != nil {
		t.Fatal(err)
	}
	recHiddenSize, err := CalcPathSize("testDir", false, true, true)
	if err != nil {
		t.Fatal(err)
	}

	if recSize != "205.1 KB" {
		t.Fatalf("Recursive noHidden: sizes do not match: %s %s", "205.1 KB", recSize)
	}
	if size != "195.3 KB" {
		t.Fatalf("NoRecursive: sizes do not match: %s %s", "195.3 KB", size)
	}
	if recHiddenSize != "293.0 KB" {
		t.Fatalf("Recursive Hidden: sizes do not match: %s %s", "293.0 KB", recHiddenSize)
	}

	errHandler(os.RemoveAll("testDir"), t)
}

func createZeroFile(path string, size int64) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close() // Игнорируем ошибку чтобы линтер не ругался
	}()
	return f.Truncate(size)
}

func errHandler(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
