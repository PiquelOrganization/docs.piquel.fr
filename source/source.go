package source

import (
	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

type Source interface {
	GetSourceType() string
	LoadFiles() utils.Files
}

func GetSource(config *config.Config) Source {
	path := "/home/piquel/Projects/OpenMC/Wiki"

	return NewFileSystemSource(path)
}
