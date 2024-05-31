package qrGenerator

import (
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	qrEnc "github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
)

var QrData = make(map[int]string, 100)

func encodeQr(text string, filename string) (string, error) {
	if len(text) > 1024 {
		return "", &ErrorLength{"Text is too long"}
	}
	qrCode, err := qrEnc.Encode(text, qrEnc.M, qrEnc.Auto)
	if err != nil {
		return "", err
	}

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return "", err
	}

	qrPath := filepath.Join("./qrs", filename)

	// create the output file
	file, err := os.Create(qrPath)
	if err != nil {
		return "", err
	}

	defer file.Close()
	// encode the barcode as png
	err = png.Encode(file, qrCode)
	if err != nil {
		return "", err
	}

	return qrPath, nil
}

func genQrs(count int) error {
	for i := 1; i <= count; i++ {
		value := uuid.New().String()
		_, err := encodeQr(value, strconv.Itoa(i)+".png")
		if err != nil {
			return err
		}
		QrData[i] = value
	}
	return nil
}

func Start() {
	err := os.Mkdir("qrs", 0755)
	if err != nil {
		log.Println(err.Error())
	}
	genQrs(100)
	ticker := time.NewTicker(15 * time.Minute)
	for _ = range ticker.C {
		err = genQrs(100)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
