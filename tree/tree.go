package tree

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ScanPathFile(path string) {
	err := filepath.Walk(path, func(itemPath string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// 计算空格数量
		depth := strings.Count(itemPath[len(path):], string(filepath.Separator))

		spaces := ""
		for i := 0; i < depth*2; i++ {
			spaces += " "
		}

		filepathBase := filepath.Base(itemPath)

		if strings.HasPrefix(filepathBase, ".") {
			//fmt.Println(filepathBase)
			//if depth != 0 {
			//	return nil
			//}
			return nil
			//return nil
		}

		fileName := info.Name()

		if info.IsDir() {
			fmt.Printf("%s|%+v \n", spaces, fileName)
		} else {
			//if strings.HasPrefix(fileName, ".") {
			//	return nil
			//}
			fmt.Printf("%s|%q\n", spaces, fileName)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ScanWalkDir(path string, depth int, maxDepth int) error {
	if depth > maxDepth {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			if strings.HasPrefix(file.Name(), ".") {
				continue
			}
			fmt.Printf("%s| %s\n", strings.Repeat("--", depth), file.Name())

			err := ScanWalkDir(filepath.Join(path, file.Name()), depth+1, maxDepth)
			if err != nil {
				return err
			}
		} else if depth > 0 {
			fmt.Printf("%s| %s\n", strings.Repeat("--", depth), file.Name())
		}
	}

	return nil
}
