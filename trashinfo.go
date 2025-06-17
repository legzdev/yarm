package yarm

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

const trashInfoFmt = `[Trash Info]
Path=%s
DeletionDate=%s
`

func WriteTrashInfo(name string, fp string) error {
	fp, err := filepath.Abs(fp)
	if err != nil {
		return err
	}

	deletionDate := time.Now().Format(time.RFC3339)

	trashInfo := fmt.Sprintf(trashInfoFmt, fp, deletionDate)
	trashInfoPath := path.Join(TrashDir, "info", name)

	file, err := os.Create(trashInfoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(trashInfo))
	if err != nil {
		return err
	}

	return nil
}
