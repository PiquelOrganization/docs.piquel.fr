package source

import (
	"fmt"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

type FileSystemSource struct{ root string }

func NewFileSystemSource(root string) *FileSystemSource {
	return &FileSystemSource{root: root}
}

func (s *FileSystemSource) GetSourceType() string {
	return "file system source"
}

func (s *FileSystemSource) LoadFiles() utils.Files {
	return s.getFilesFromDir(s.root)
}

func (s *FileSystemSource) getFilesFromDir(path string) utils.Files {
	files := utils.Files{}

	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		name := entry.Name()
		if entry.IsDir() {
			files = append(files, s.getFilesFromDir(fmt.Sprintf("%s/%s", path, name))...)
			continue
		}

		if !strings.HasSuffix(name, ".md") {
			continue
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", path, name))
		if err != nil {
			panic(err)
		}

		filePath := strings.Replace(file.Name(), s.root, "", 1)
		filePath = strings.Replace(filePath, ".md", "", 1)
		filePath = strings.Trim(filePath, "/")
		filePath = fmt.Sprintf("/%s", filePath)

		fileData, err := os.ReadFile(file.Name())
        if err != nil {
            panic(err)
        }

		docsFile := utils.File{
			Path: filePath,
			Data: fileData,
		}

		files = append(files, docsFile)
	}

	return files
}
