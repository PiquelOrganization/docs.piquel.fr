package source

import (
	"errors"
	"fmt"
	"log"
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
	pages := s.getFilesFromDir(s.config.DataPath, "md")
	commonFolder := fmt.Sprintf("%s/.common", s.config.DataPath)

	if file, err := os.Stat(commonFolder); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("[Source] No common folder")
			return utils.SourceDocs{Pages: pages}
		}
		panic(err)
	} else {
		if !file.IsDir() {
			log.Printf("[Source] No common folder")
			return utils.SourceDocs{Pages: pages}
		}
	}

	includes := s.getFilesFromDir(fmt.Sprintf("%s/includes", commonFolder), ".md")
	styles := s.getFilesFromDir(fmt.Sprintf("%s/styles", commonFolder), ".css")
	assets := s.getFilesFromDir(fmt.Sprintf("%s/assets", commonFolder), "")

	return utils.SourceDocs{
		Pages:    pages,
		Includes: includes,
		Styles:   styles,
		Assets:   assets,
	}
}

func (s *RealSource) getFilesFromDir(path, ext string) utils.Files {
	pages := utils.Files{}

	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		name := entry.Name()

		if strings.HasPrefix(name, ".") {
			continue
		}

		if entry.IsDir() {
			maps.Copy(pages, s.getFilesFromDir(fmt.Sprintf("%s/%s", path, name), ext))
			continue
		}

		if ext != "" && !strings.HasSuffix(name, ext) {
			continue
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", path, name))
		if err != nil {
			panic(err)
		}

		filePath := strings.Replace(file.Name(), s.config.DataPath, "", 1)
		if ext != "" {
			filePath = strings.ReplaceAll(filePath, ext, "")
		}
		filePath = strings.Trim(filePath, "/")
		filePath = strings.Trim(filePath, ".")
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
