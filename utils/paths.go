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

func ValidatePath(path string) bool {
	if strings.Contains(path, "..") {
		return false
	} else if strings.Contains(path, "~") {
		return false
	}
	return true
}

func FormatLocalPathString(path, ext string) string {
	trim := strings.Trim(path, "/")
	trim = strings.TrimSuffix(trim, ext)
	return fmt.Sprintf("/%s", trim)
}

func FormatLocalPath(path []byte, ext string) []byte {
	trim := bytes.Trim(path, "/")
	trim = bytes.TrimSuffix(trim, []byte(ext))
	return slices.Concat([]byte("/"), trim)
}
