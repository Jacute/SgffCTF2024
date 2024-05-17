package main

import (
	"log"
	"os"
	"time"
)

func runCleaner(dirPath string) {
	ticker := time.NewTicker(15 * time.Minute)
	for _ = range ticker.C {
		err := os.RemoveAll(dirPath)
		os.Mkdir(dirPath, 0755)
		if err != nil {
			log.Printf("Directory %s wasn't cleaned", dirPath)
		}
	}
}
