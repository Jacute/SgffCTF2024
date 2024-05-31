package main

import (
	"math/rand"
	"os"
	"strconv"
)

const FLAG = "SgffCTF{ctf_1n_ku1by5h3v_c1ty}"

func writeFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0764)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(data)

	return nil
}

func main() {
	os.Mkdir("files", 0764)
	c := 0
	for i := 1; i <= 1000; i++ {
		filename := "files/" + strconv.Itoa(i) + ".txt"
		if i > 659 && c < len(FLAG) {
			writeFile(filename, []byte{FLAG[c]})
			c++
		} else {
			writeFile(filename, []byte{byte(33 + rand.Intn(90))})
		}
	}
}
