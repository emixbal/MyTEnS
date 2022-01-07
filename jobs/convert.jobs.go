package jobs

import (
	"MyTensApp/helpers"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Convert(filePath, fileTypeFinalPayload, fileLocFinal string) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return err
	}

	fileTypeFinal := "txt"
	switch fileTypeFinalPayload {
	case "text":
		fileTypeFinal = "txt"
	case "json":
		fileTypeFinal = "json"
	default:
		return errors.New("type is only accept \"text\" or \"json\"")
	}

	if fileLocFinal != "same" {
		if _, err := os.Stat(fileLocFinal); os.IsNotExist(err) {
			return errors.New("invalid destination")
		}
	} else {
		fileLocFinal = "/var/log/"
	}

	//extract file name
	fileNameExt := filepath.Base(filePath)
	fileName := strings.TrimSuffix(fileNameExt, path.Ext(fileNameExt))

	//define new file path
	newFIlePath := fileLocFinal + "/" + fileName + "." + fileTypeFinal

	// start copy file and create new one with diferent file type
	_, errCopy := helpers.Copy(filePath, newFIlePath)
	if errCopy != nil {
		return errCopy
	}

	return nil
}
