package files

import (
	"os"
	"path/filepath"

	"github.com/PaulioRandall/go-trackerr"
)

var (
	ErrReadingFileList = trackerr.New("Could not read files in directory")
	ErrCreatingAbsPath = trackerr.New("Could not create absolute file path")
)

type ReadOnlyFile struct {
	Name    string
	Path    string
	AbsPath string
	IsDir   bool
}

func ListFilesInDir(dir string) ([]ReadOnlyFile, error) {
	children, e := os.ReadDir(dir)
	if e != nil {
		return nil, ErrReadingFileList.CausedBy(e)
	}

	files := make([]ReadOnlyFile, len(children)+1)

	files[0], e = createParentDirEntry(dir)
	if e != nil {
		return nil, ErrReadingFileList.CausedBy(e)
	}

	for i := 0; i < len(children); i++ {
		files[i+1], e = mapToReadOnlyFile(dir, children[i])
		if e != nil {
			return nil, ErrReadingFileList.CausedBy(e)
		}
	}

	return files, e
}

func mapToReadOnlyFile(parent string, file os.DirEntry) (ReadOnlyFile, error) {
	var (
		rof  ReadOnlyFile
		zero ReadOnlyFile
		e    error
	)

	rof.Name = file.Name()
	rof.Path = filepath.Join(parent, rof.Name)
	rof.IsDir = true

	rof.AbsPath, e = ToAbsPath(rof.Path)
	if e != nil {
		return zero, e
	}

	rof.IsDir, e = isFileDir(rof.AbsPath)
	if e != nil {
		return zero, e
	}

	return rof, nil
}

func createParentDirEntry(currDir string) (ReadOnlyFile, error) {
	var zero ReadOnlyFile

	absPath, e := ToAbsPath(currDir)
	if e != nil {
		return zero, e
	}

	rof := ReadOnlyFile{
		Name:  "..",
		IsDir: true,
		Path:  filepath.Dir(absPath),
	}

	rof.AbsPath, e = ToAbsPath(rof.Path)
	if e != nil {
		return zero, e
	}

	return rof, nil
}

func ToAbsPath(path string) (string, error) {
	absPath, e := filepath.Abs(path)
	if e != nil {
		return "", ErrCreatingAbsPath.CausedBy(e)
	}
	return absPath, nil
}

func isFileDir(path string) (bool, error) {
	info, e := os.Stat(path)
	if e != nil {
		return false, ErrReadingFileList.CausedBy(e)
	}
	return info.IsDir(), nil
}
