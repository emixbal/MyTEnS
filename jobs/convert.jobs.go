package jobs

import (
	"MyTEnsApp/helpers"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Convert(filePath, fileTypeFinalPayload, filePathGiven string) error {
	// check is log exist?
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return err
	}

	//define new extension
	fileTypeFinal := "txt"
	switch fileTypeFinalPayload {
	case "text":
		fileTypeFinal = ".txt"
	case "json":
		fileTypeFinal = ".json"
	default:
		return errors.New("type is only accept \"text\" or \"json\"")
	}

	newFilePath := ""
	if filePathGiven != "same" { // if opsi tidak dikosongkan
		fileBasePathFinal, _ := path.Split(filePathGiven)
		fileNameGiven := filepath.Base(filePathGiven)
		fileTypeGiven := filepath.Ext(fileNameGiven)

		if _, err := os.Stat(fileBasePathFinal); os.IsNotExist(err) {
			return errors.New("invalid option destination")
		}

		if fileTypeGiven != fileTypeFinal {
			return errors.New("given file name extension must be same with flag type")
		}
		newFilePath = filePathGiven
	} else { // if opsi dikosongkan
		fileBasePathFinal, _ := path.Split(filePath)
		fileNameExt := filepath.Base(filePath)
		fileBaseName := strings.TrimSuffix(fileNameExt, filepath.Ext(fileNameExt))
		newFilePath = fileBasePathFinal + fileBaseName + fileTypeFinal
	}

	// start copy file and create new one with diferent file type
	_, errCopy := helpers.Copy(filePath, newFilePath)
	if errCopy != nil {
		return errCopy
	}

	return nil
}
