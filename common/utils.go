package common

import (
	"os"
	"path/filepath"
	"strings"
)

func WalkTemplates(templateDir, suffix string) (pages []string, err error) {
	pages = make([]string, 0, 30)
	suffix = strings.ToLower(suffix)
	err = filepath.Walk(templateDir, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			pages = append(pages, filename)
		}
		return nil
	})
	return pages, err
}
