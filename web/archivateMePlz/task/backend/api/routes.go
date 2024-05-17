package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func getArchive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	name := r.URL.Query().Get("file")
	if name == "" {
		http.Error(w, "File name not provided", http.StatusBadRequest)
		return
	}

	path := filepath.Join("archives", name)
	f, err := os.Open(path)
	if err != nil {
		http.Error(w, "Unable to open archive", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	http.ServeContent(w, r, name, time.Time{}, f)
}

func createArchive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filesToArchivate := make([]string, 0, 10)
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}
		defer part.Close()

		fileName := part.FileName()
		if fileName == "" {
			continue
		}

		data, err := io.ReadAll(part)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filePath, err := saveFile(fileName, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filesToArchivate = append(filesToArchivate, filePath)
	}
	zipName, err := archivateFiles(filesToArchivate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		archive := Archive{
			Name: zipName,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(archive)
	}

}
