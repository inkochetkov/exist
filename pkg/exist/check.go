package exist

import "path"

const (
	empty = ""
)

// InitDirFile - check dirictory and file
func InitDirFile(pathDir, nameFile string) (string, error) {

	err := InitDir(pathDir)
	if err != nil {
		return empty, err
	}

	pathFile := path.Join(pathDir, nameFile)

	err = InitFile(pathFile)
	if err != nil {
		return empty, err
	}

	return pathFile, nil

}
