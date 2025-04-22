package source

import (
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
	files := utils.SourceDocs{}
	files.Pages = s.getFilesFromDir(s.config.DataPath, "md", s.config.DataPath, true)
	commonFolder := fmt.Sprintf("%s/.common", s.config.DataPath)

	if !isDir(commonFolder) {
		log.Printf("[Source] No common folder")
		return files
	}

	includesDir := fmt.Sprintf("%s/includes", commonFolder)
	if isDir(includesDir) {
		files.Includes = s.getFilesFromDir(includesDir, ".md", includesDir, true)
	}

	stylesDir := fmt.Sprintf("%s/styles", commonFolder)
	if isDir(stylesDir) {
		files.Styles = s.getFilesFromDir(stylesDir, ".css", stylesDir, false)
	}

	assetsDir := fmt.Sprintf("%s/assets", commonFolder)
	if isDir(assetsDir) {
		files.Assets = s.getFilesFromDir(assetsDir, "", assetsDir, false)
	}

	return files
}

func (s *RealSource) getFilesFromDir(path, ext, root string, removeExt bool) utils.Files {
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
			maps.Copy(pages, s.getFilesFromDir(fmt.Sprintf("%s/%s", path, name), ext, root, removeExt))
			continue
		}

		if ext != "" && !strings.HasSuffix(name, ext) {
			continue
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", path, name))
		if err != nil {
			panic(err)
		}

		filePath := strings.Replace(file.Name(), root, "", 1)
		if ext != ""  && removeExt{
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

func isDir(path string) bool {
	if file, err := os.Stat(path); err != nil {
		return false
	} else {
		return file.IsDir()
	}
}
