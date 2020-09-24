package util

import "os"

func GetFileSize(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}
	total := fi.Size()
	err = f.Close()
	if err != nil {
		return 0, err
	}
	return total, nil
}
