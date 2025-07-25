package source

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/piquel-fr/piquel-docs/config"
	"github.com/piquel-fr/piquel-docs/git"
	"github.com/piquel-fr/piquel-docs/utils"
	"gopkg.in/yaml.v3"
)

type Source interface {
	Update() error
	GetAllMarkdown() ([]string, error)
	LoadRouteFile(route string) ([]byte, error)
	LoadInclude(path string) ([]byte, error)
	GetAssetsPath() string
}

type GitSource struct {
	dataPath, repository string
	docsConfig           *config.DocsConfig
}

func NewGitSource(config *config.Config) Source {
	return &GitSource{
		dataPath:   config.Envs.DataPath,
		repository: config.Envs.Repository,
		docsConfig: &config.Config,
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

	configData, err := os.ReadFile(fmt.Sprintf("%s/config.yml", s.dataPath))
	if errors.Is(err, os.ErrNotExist) {
		configData, err = os.ReadFile(fmt.Sprintf("%s/config.yaml", s.dataPath))
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("[Source] Could not find configuration file in repository")
			return nil
		}
	}
	if err != nil {
		return err
	}

	s.docsConfig.Lock()
	if err = yaml.Unmarshal(configData, &s.docsConfig); err != nil {
		return err
	}

	s.docsConfig.HomePage = utils.FormatLocalPathString(s.docsConfig.HomePage, ".md")
	s.docsConfig.Root = utils.FormatLocalPathString(s.docsConfig.Root, ".md")

	s.docsConfig.Unlock()

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
		name = utils.FormatLocalPathString(name, ext)

		files = append(files, name)
	}

	return files, nil
}

func (s *GitSource) LoadRouteFile(route string) ([]byte, error) {
	err := utils.ValidatePath(route)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s%s.md", s.dataPath, route)

	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	return os.ReadFile(fileName)
}

func (s *GitSource) LoadInclude(path string) ([]byte, error) {
	err := utils.ValidatePath(path)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s/.common/includes%s.md", s.dataPath, utils.FormatLocalPathString(path, ".md"))

	if _, err := os.Stat(fileName); err != nil {
		return nil, err
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
