package yarm

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"
)

const trashInfoFmt = `[Trash Info]
Path=%s
DeletionDate=%s
`

func CreateTrashInfoFile(name string) (info *os.File, baseName string, err error) {
	baseName = path.Base(name)

	for {
		infoPath := path.Join(TrashDir, "info", baseName+".trashinfo")

		file, err := os.OpenFile(infoPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if err == nil {
			return file, baseName, nil
		}

		if errors.Is(err, fs.ErrExist) {
			baseName = path.Base(name) + "." + GenerateRandomID()
			continue
		}

		return nil, "", err
	}
}

func WriteTrashInfo(file *os.File, path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	deletionDate := time.Now().Format(time.RFC3339)
	trashInfo := fmt.Sprintf(trashInfoFmt, path, deletionDate)

	_, err = file.Write([]byte(trashInfo))
	if err != nil {
		return err
	}

	return nil
}
