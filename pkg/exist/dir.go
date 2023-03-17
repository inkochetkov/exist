package exist

import "os"

// InitDir  initialization dir
func InitDir(path string) error {

	if check := CheckDir(path); check {
		return nil
	}

	err := CreateDir(path)
	if err != nil {
		return err
	}

	return nil
}

// CheckDir - dir existence check
func CheckDir(path string) bool {

	fs, err := os.Lstat(path)
	if err != nil || !fs.IsDir() {
		return false
	}

	return true
}

// CreateDir - create dir
func CreateDir(path string) error {

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
