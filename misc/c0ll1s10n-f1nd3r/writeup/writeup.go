package main

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"runtime"
)

// worker функция, которая проверяет значения и отправляет результат в канал, если коллизия найдена
func worker(targetCRC uint32, start, step uint64, resultChan chan<- []byte) {
	for i := start; i < start+step; i++ {
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, i)
		if crc32.ChecksumIEEE(bytes) == targetCRC {
			resultChan <- bytes
			return
		}
	}
}

func goroutineStart(targetCRC uint32) []byte {
	workersCount := uint64(runtime.NumCPU())
	step := 0xFFFFFFFFFFFFFFFF / workersCount

	resultChan := make(chan []byte)

	for i := uint64(0); i < workersCount; i++ {
		go worker(targetCRC, i*step, step, resultChan)
	}
	return <-resultChan
}

func main() {
	targetCRC := uint32(0x8eff10ee)

	collisionData := goroutineStart(targetCRC)
	fmt.Printf("Data: %x\n", collisionData)
	fmt.Printf("CRC-32: %08X\n", crc32.ChecksumIEEE(collisionData))
}
