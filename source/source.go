package source

import (
	"fmt"
	"maps"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

type Source interface {
	LoadFiles() utils.SourceDocs
}

type RealSource struct {
	config *config.Config
}

func NewSource(config *config.Config) Source {
	return &RealSource{config}
}

func (s *RealSource) LoadFiles() utils.SourceDocs {
	//s.getFilesFromDir(s.config.DataPath)
	return utils.SourceDocs{}
}

func (s *RealSource) getFilesFromDir(path string) utils.Pages {
	pages := utils.Pages{}

	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		name := entry.Name()
		if entry.IsDir() {
			maps.Copy(pages, s.getFilesFromDir(fmt.Sprintf("%s/%s", path, name)))
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
		filePath = strings.ToLower(filePath)

		fileData, err := os.ReadFile(file.Name())
		if err != nil {
			panic(err)
		}

		pages[filePath] = fileData
	}

	return pages
}
