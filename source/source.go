package source

import (
	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/git"
)

type Source interface {
	Update() error
	GetAllFiles() []string
	LoadFile(path string) []byte
	LoadInclude(path string) []byte
	GetAssetsPath() string
}

type GitSource struct {
	config struct {
		dataPath, repository string
	}
}

func NewGitSource(config *config.Config) Source {
	return &GitSource{config: struct {
		dataPath   string
		repository string
	}{
		config.DataPath,
		config.Repository,
	}}
}

func (s *GitSource) Update() error {
	if err := git.Status(s.config.dataPath); err == nil {
		err = git.Pull(s.config.repository)
		if err != nil {
			return err
		}
	} else {
		err := git.Clone(s.config.repository, s.config.dataPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GitSource) GetAllFiles() []string {
	// TODO
	return []string{}
}

func (s *GitSource) LoadFile(path string) []byte {
	// TODO
	return []byte{}
}

func (s *GitSource) LoadInclude(path string) []byte {
	// TODO
	return []byte{}
}

func (s *GitSource) GetAssetsPath() string {
	// TODO
	return ""
}

// func (s *GitSource) LoadFiles() utils.SourceDocs {
// 	files := utils.SourceDocs{}
// 	files.Pages = s.getFilesFromDir(s.config.DataPath, s.config.DataPath, ".md")
// 	commonFolder := fmt.Sprintf("%s/.common", s.config.DataPath)
//
// 	if !isDir(commonFolder) {
// 		log.Printf("[Source] No common folder")
// 		return files
// 	}
//
// 	includesDir := fmt.Sprintf("%s/includes", commonFolder)
// 	if isDir(includesDir) {
// 		files.Includes = s.getFilesFromDir(includesDir, includesDir, ".md")
// 	}
//
// 	stylesDir := fmt.Sprintf("%s/styles", commonFolder)
// 	if isDir(stylesDir) {
// 		files.Styles = s.getFilesFromDir(stylesDir, stylesDir, ".css")
// 	}
//
// 	assetsDir := fmt.Sprintf("%s/assets", commonFolder)
// 	if isDir(assetsDir) {
// 		files.Assets = s.getFilesFromDir(assetsDir, assetsDir, "")
// 	}
//
// 	return files
// }
//
// func (s *GitSource) getFilesFromDir(root, path, ext string) utils.Files {
// 	pages := utils.Files{}
//
// 	dir, err := os.ReadDir(path)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	for _, entry := range dir {
// 		name := entry.Name()
//
// 		if strings.HasPrefix(name, ".") {
// 			continue
// 		}
//
// 		newPath := fmt.Sprintf("%s/%s", path, name)
// 		if entry.IsDir() {
// 			maps.Copy(pages, s.getFilesFromDir(root, newPath, ext))
// 			continue
// 		}
//
// 		if ext != "" && !strings.HasSuffix(name, ext) {
// 			continue
// 		}
//
// 		file, err := os.Open(newPath)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		filePath := strings.Replace(file.Name(), root, "", 1)
// 		filePath = strings.Trim(filePath, "/")
// 		filePath = strings.ToLower(filePath)
//
// 		fileData, err := os.ReadFile(file.Name())
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		pages[filePath] = fileData
// 	}
//
// 	return pages
// }
