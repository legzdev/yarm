package yarm

import (
	"os"
	"path"
)

func MoveToTrash(target string) error {
	err := CheckTarget(target)
	if err != nil {
		return err
	}

	inTrashDir := InTrashDir(target)
	if inTrashDir {
		ok, err := confirmf("Permanently remove '%s'?", target)
		if err != nil {
			return err
		}

		if !ok {
			verbosef("skipping '%s'", target)
			return nil
		}

		// TODO: print removed files
		return os.RemoveAll(target)
	}

	trashinfoFile, baseName, err := CreateTrashInfoFile(target)
	if err != nil {
		return err
	}
	defer trashinfoFile.Close()

	trashPath := path.Join(TrashDir, "files", baseName)

	err = moveToTrash(target, trashPath, trashinfoFile)
	if err != nil {
		return err
	}

	verbosef("trashed '%s' => '%s'\n", target, trashPath)
	return nil
}

func moveToTrash(target, trashPath string, trashInfo *os.File) error {
	if FlagDryRun {
		return nil
	}

	err := os.Rename(target, trashPath)
	if err != nil {
		return err
	}

	err = WriteTrashInfo(trashInfo, target)
	if err != nil {
		return err
	}

	return nil
}

func CheckTarget(target string) error {
	targetInfo, err := os.Stat(target)
	if err != nil {
		return err
	}

	if !targetInfo.IsDir() || FlagRecursive {
		return nil
	}

	if !FlagDir {
		return &ErrCannotRemove{Msg: "Is a directory", Target: target}
	}

	isEmpty, err := IsDirEmpty(target)
	if err != nil {
		return err
	}

	if isEmpty {
		return nil
	}

	return &ErrCannotRemove{Msg: "Directory not empty", Target: target}
}
