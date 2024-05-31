package main

import (
	"fmt"
	"log"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"

	qrDec "github.com/tuotoo/qrcode"

	"github.com/go-resty/resty/v2"
)

const host = "http://localhost:1338"

func decodeQr(qrPath string) (string, error) {
	file, err := os.Open(qrPath)
	if err != nil {
		log.Println("Error opening the file:", err)
		return "", err
	}
	defer file.Close()

	qrCode, err := qrDec.Decode(file)
	if err != nil {
		log.Println("Error when decoding the QR code:", err)
		return "", err
	}

	return qrCode.Content, nil
}

func writeFile(filename string, bytes []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0764)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(bytes)

	return nil
}

func main() {
	u, _ := url.Parse(host)

	jar, _ := cookiejar.New(nil)
	client := resty.New()
	client.SetCookieJar(jar)

	_, err := client.R().Get(host)
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 100; i++ {
		filename := strconv.Itoa(i) + ".png"
		res, _ := client.R().
			Get(host + "/qrs/" + filename)
		qrPath := "qrs/" + filename
		writeFile(qrPath, res.Body())

		decoded, err := decodeQr(qrPath)
		if err != nil {
			panic(err)
		}

		res, err = client.R().
			SetHeader("Content-Type", "text/plain").
			SetFormData(map[string]string{
				"answer": decoded,
			}).
			Post(host)
		if err != nil {
			panic(err)
		}
		fmt.Println("GET response: ", res)
	}

	res, err := client.R().Get(host)
	fmt.Println(string(res.Body()))
	fmt.Println("Cookie: ", jar.Cookies(u))
}
