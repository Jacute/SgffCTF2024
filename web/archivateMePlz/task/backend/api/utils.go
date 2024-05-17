package api

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func saveFile(fileName string, data []byte) (string, error) {
	filePath := filepath.Join("uploads", fileName)

	file, err := os.OpenFile("uploads/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func deleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func archivateFiles(files []string) (string, error) {
	zipName := uuid.New().String() + ".zip"
	zipPath := filepath.Join("archives", zipName)

	zipFile, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return "", err
		}
		defer file.Close()

		zipFile, err := zipWriter.Create(filepath.Base(filename))
		if err != nil {
			return "", err
		}

		_, err = io.Copy(zipFile, file)
		if err != nil {
			return "", err
		}
		deleteFile(filename)
	}
	return zipName, nil
}
