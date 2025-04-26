package source

import (
	"fmt"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/git"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

type Source interface {
	Update() error
	GetAllMarkdown() ([]string, error)
	LoadFile(path string) ([]byte, error)
	LoadInclude(path string) ([]byte, error)
	GetAssetsPath() string
}

type GitSource struct {
	dataPath, repository string
}

func NewGitSource(config *config.Config) Source {
	return &GitSource{
		dataPath:   config.DataPath,
		repository: config.Repository,
	}
}

func (s *GitSource) Update() error {
	if err := git.Status(s.dataPath); err == nil {
		err = git.Pull(s.dataPath)
		if err != nil {
			return err
		}
	} else {
		err := git.Clone(s.repository, s.dataPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GitSource) GetAllMarkdown() ([]string, error) {
	return s.readDir(s.dataPath, ".md")
}

func (s *GitSource) readDir(path, ext string) ([]string, error) {
	files := []string{}

	dir, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	for _, entry := range dir {
		name := entry.Name()

		if strings.HasPrefix(name, ".") {
			continue
		}

		newPath := fmt.Sprintf("%s/%s", path, name)
		if entry.IsDir() {
			newFiles, err := s.readDir(newPath, ext)
			if err != nil {
				return files, err
			}

			files = append(files, newFiles...)
			continue
		}

		if ext != "" && !strings.HasSuffix(name, ext) {
			continue
		}

		name = strings.Replace(newPath, s.dataPath, "", 1)
		name = strings.Trim(name, "/")
		name = strings.ReplaceAll(name, ext, "")

		files = append(files, name)
	}

	return files, nil
}

func (s *GitSource) LoadFile(path string) ([]byte, error) {
	fileName := fmt.Sprintf("%s/%s.md", s.dataPath, strings.Trim(path, "/"))

	if _, err := os.Stat(fileName); err != nil {
		return []byte{}, err
	}

	return os.ReadFile(fileName)
}

func (s *GitSource) LoadInclude(path string) ([]byte, error) {
	fileName := fmt.Sprintf("%s/.common/includes/%s.md", s.dataPath, strings.Trim(path, "/"))

	if _, err := os.Stat(fileName); err != nil {
		return []byte{}, err
	}

	return os.ReadFile(fileName)
}

func (s *GitSource) GetAssetsPath() string {
	assetsDir := fmt.Sprintf("%s/.common/assets", s.dataPath)
	if utils.IsDir(assetsDir) {
		return assetsDir
	}
	return ""
}
