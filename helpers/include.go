package helpers

import (
	"fmt"
	"path/filepath"
)

func Include(path string) []string {
	files, err := filepath.Glob("views/templates/*.html")
	if err != nil {
		fmt.Println(err)
	}
	path_files, _ := filepath.Glob("views/" + path + "/*.html")
	for _, file := range path_files {
		files = append(files, file)
	}
	return files
}
