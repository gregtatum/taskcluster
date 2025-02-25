package d2g

import (
	"github.com/taskcluster/shell"
	"github.com/taskcluster/taskcluster/v74/tools/d2g/genericworker"
)

func (ndi *NamedDockerImage) FileMounts(tool string) ([]genericworker.FileMount, error) {
	return []genericworker.FileMount{}, nil
}

func (ndi *NamedDockerImage) String(tool string) (string, error) {
	return shell.Escape(ndi.Name), nil
}

func (ndi *NamedDockerImage) LoadCommands(tool string) []string {
	return []string{}
}
