package source

import (
	"fmt"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

type Source interface {
	LoadFiles() utils.Files
}

type RealSource struct {
	config *config.Config
}

func NewSource(config *config.Config) Source {
	return &RealSource{config}
}

func (s *RealSource) LoadFiles() utils.Files {
	return s.getFilesFromDir(s.config.DataPath)
}

func (s *RealSource) getFilesFromDir(path string) utils.Files {
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

		filePath := strings.Replace(file.Name(), s.config.DataPath, "", 1)
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
