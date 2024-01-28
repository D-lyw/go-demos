package tree

import "testing"

func TestScanPathFile(t *testing.T) {
	testPath := "/Users/d-lyw/GoProjects/go-demos"
	ScanPathFile(testPath)
}

func TestScanWalkDir(t *testing.T) {
	testPath := "/Users/d-lyw/GoProjects/go-demos"
	err := ScanWalkDir(testPath, 0)
	if err != nil {
		return
	}
}
