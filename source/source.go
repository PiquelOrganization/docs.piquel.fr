package source

import (
	"fmt"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

func LoadFiles(config *config.Config) utils.Files {
	return getFilesFromDir(config, config.DataPath)
}

func getFilesFromDir(config *config.Config, path string) utils.Files {
	files := utils.Files{}

	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		name := entry.Name()
		if entry.IsDir() {
			files = append(files, getFilesFromDir(config, fmt.Sprintf("%s/%s", path, name))...)
			continue
		}

		if !strings.HasSuffix(name, ".md") {
			continue
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", path, name))
		if err != nil {
			panic(err)
		}

		filePath := strings.Replace(file.Name(), config.DataPath, "", 1)
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
