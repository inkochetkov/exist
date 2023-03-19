package exist

// InitDirFile - check dirictory and file
func InitDirFile(path, nameFile string) error {

	err := InitDir(path)
	if err != nil {
		return err
	}

	err = InitFile(nameFile)
	if err != nil {
		return err
	}

	return nil

}
