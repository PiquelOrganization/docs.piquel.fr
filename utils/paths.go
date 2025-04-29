package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func IsDir(path string) bool {
	if file, err := os.Stat(path); err != nil {
		return false
	} else {
		return file.IsDir()
	}
}

func FormatPathSlashesString(path string) string {
	trim := strings.Trim(path, "/")
	if strings.HasPrefix(trim, "http") {
		return trim
	}
	return fmt.Sprintf("/%s", trim)
}

func FormatPathSlashes(path []byte) []byte {
	trim := bytes.Trim(path, "/")
	if bytes.HasPrefix(trim, []byte("http")) {
		return trim
	}
	return append([]byte("/"), trim...)
}
