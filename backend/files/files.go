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
	ErrCheckingStatus  = trackerr.New("Could not check status of file")
)

func ToAbsPath(path string) (string, error) {
	absPath, e := filepath.Abs(path)
	if e != nil {
		return "", ErrCreatingAbsPath.CausedBy(e)
	}
	return absPath, nil
}

func DoesFileExist(file string) (bool, error) {
	if _, e := os.Stat(file); e == nil {
		return true, nil
	} else if trackerr.Is(e, os.ErrNotExist) {
		return false, nil
	} else {
		return false, ErrCheckingStatus.BecauseOf(e, "For '%s'", file)
	}
}

func ListFilesInDir(dir string) ([]ReadOnlyFile, error) {
	children, e := os.ReadDir(dir)
	if e != nil {
		return errReadingFileList(e, dir)
	}

	files := []ReadOnlyFile{}

	if dir != "/" {
		parentDir, e := createParentDirEntry(dir)
		if e != nil {
			return errReadingFileList(e, dir)
		}

		files = append(files, parentDir)
	}

	for i := 0; i < len(children); i++ {
		child, e := mapToReadOnlyFile(dir, children[i])
		if e != nil {
			return errReadingFileList(e, dir)
		}

		files = append(files, child)
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

func isFileDir(path string) (bool, error) {
	info, e := os.Stat(path)
	if e != nil {
		return false, ErrReadingFileStat.CausedBy(e)
	}
	return info.IsDir(), nil
}
