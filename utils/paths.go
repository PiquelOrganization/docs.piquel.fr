package utils

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func IsDir(path string) bool {
	if file, err := os.Stat(path); err != nil {
		return false
	} else {
		return file.IsDir()
	}
}

func FormatLocalPathString(path string) string {
	trim := strings.Trim(path, "/")
	trim = strings.TrimSuffix(trim, ".md")
	return fmt.Sprintf("/%s", trim)
}

func FormatLocalPath(path []byte) []byte {
	trim := bytes.Trim(path, "/")
	trim = bytes.TrimSuffix(trim, []byte(".md"))
	return slices.Concat([]byte("/"), trim)
}
