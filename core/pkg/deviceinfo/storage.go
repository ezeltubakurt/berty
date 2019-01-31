package deviceinfo

import (
	"errors"
	"os"
)

var storagePath = ""

func SetStoragePath(path string) error {
	stat, err := os.Stat(path)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0700); err != nil {
			return err
		}
	}

	if stat.IsDir() == false {
		return errors.New("storage path is not a directory: not a dir")
	}

	storagePath = path
	return nil
}

func GetStoragePath() string {
	return storagePath
}
