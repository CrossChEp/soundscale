package repository

import (
	"fmt"
	"os"
	"photo_service/pkg/logger"
)

func Upload(path string, file []byte) error {
	if err := os.WriteFile(path, file, 0644); err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload file. Details: %v", err))
		return err
	}
	return nil
}

func Download(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("(Download)Error: couldn't read file. Details: %v", err))
		return nil, err
	}
	return file, nil
}
