package exist

import "path"

// InitDirFile - check dirictory and file
func InitDirFile(pathDir, nameFile string) error {

	err := InitDir(pathDir)
	if err != nil {
		return err
	}

	pathFile := path.Join(pathDir, nameFile)

	err = InitFile(pathFile)
	if err != nil {
		return err
	}

	return nil

}
