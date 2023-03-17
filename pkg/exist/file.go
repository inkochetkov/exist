package exist

import "os"

// InitFile  initialization file
func InitFile(path string) error {

	if check := CheckFile(path); check {
		return nil
	}

	file, err := CreateFile(path)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// CheckFile- file existence check
func CheckFile(path string) bool {

	fs, err := os.Lstat(path)
	if err != nil ||
		os.IsNotExist(err) ||
		fs.IsDir() {
		return false
	}

	return true
}

// CreateFiele - create file
func CreateFile(path string) (*os.File, error) {

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// GetFile - get file
func GetFile(path string) (*os.File, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
