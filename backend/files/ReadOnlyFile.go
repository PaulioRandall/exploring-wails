package files

import (
	"strings"
)

type ReadOnlyFile struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	AbsPath string `json:"absPath"`
	IsDir   bool   `json:"isDir"`
}

func (rof ReadOnlyFile) String() string {
	sb := strings.Builder{}
	sb.WriteString("{\n")

	sb.WriteString("\tName: ")
	sb.WriteString(rof.Name)

	sb.WriteString(",\n\tPath: ")
	sb.WriteString(rof.Path)

	sb.WriteString(",\n\tAbsPath: ")
	sb.WriteString(rof.AbsPath)

	sb.WriteString(",\n\tIsDir: ")
	if rof.IsDir {
		sb.WriteString("true")
	} else {
		sb.WriteString("false")
	}

	sb.WriteString("\n}")
	return sb.String()
}
