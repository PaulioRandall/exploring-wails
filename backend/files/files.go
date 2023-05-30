package files

import (
	"os"
	"path/filepath"

	"github.com/PaulioRandall/go-trackerr"
)

var (
	ErrReadingFileList = trackerr.New("Could not read files in directory")
	ErrReadingFileStat = trackerr.New("Could not read file status (os.Stat)")
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
		return errReadingFileList(e, dir)
	}

	files := make([]ReadOnlyFile, len(children)+1)

	i := 0
	if dir != "/" {
		files[i], e = createParentDirEntry(dir)
		i++
		if e != nil {
			return errReadingFileList(e, dir)
		}
	}

	for ; i < len(children); i++ {
		files[i], e = mapToReadOnlyFile(dir, children[i])
		if e != nil {
			return errReadingFileList(e, dir)
		}
	}

	return files, nil
}

func errReadingFileList(e error, dir string) ([]ReadOnlyFile, error) {
	return nil, ErrReadingFileList.BecauseOf(e, "For '%s'", dir)
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
		return false, ErrReadingFileStat.CausedBy(e)
	}
	return info.IsDir(), nil
}
