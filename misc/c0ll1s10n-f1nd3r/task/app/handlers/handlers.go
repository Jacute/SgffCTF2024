package handlers

import (
	"bytes"
	"hash/crc32"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func PostCheck(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Couldn't upload file")
		return
	}

	dst := filepath.Join("./uploads", uuid.New().String())

	if err = c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusInternalServerError, "Couldn't save file")
		return
	}

	fileBytes, err := os.ReadFile(dst)
	os.Remove(dst)
	if err != nil {
		c.String(http.StatusInternalServerError, "Couldn't read file")
		return
	}

	if crc32.ChecksumIEEE(fileBytes) != targetCRC32 {
		c.String(http.StatusBadRequest, "Invalid CRC-32")
		return
	}

	if bytes.Equal(fileBytes, logo) {
		c.String(http.StatusOK, "Yeah! This is \"Продам гараж за флаги\" logo!")
		return
	}

	c.String(http.StatusOK, flag)
}
