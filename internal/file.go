package internal

import (
	"os"
	"path/filepath"
	"time"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func FileAge(path string) (time.Duration, error) {
	if mtime, err := Mtime(path); err != nil {
		return time.Duration(0), err
	} else {
		return time.Since(mtime), nil
	}
}

func Mtime(path string) (time.Time, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}

	return stat.ModTime(), nil
}

func WriteFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(data)
	return err
}
