package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var characterSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func main() {
	start := time.Now()
	seed := uint64(3)
	prefix := []byte("Bola is a boy")
	var str []byte
	hashes := 0
	for {
		hashes++
		seed, str = RandomString(20, seed)
		if Hash(append(prefix, str...), 3) {
			fmt.Println(string(str))
			break
		}
	}
	end := time.Now()
	time := end.Sub(start).Seconds()
	fmt.Println("total time:", time)
	fmt.Println("No of hashes:", hashes)
	fmt.Printf("processed/sec: %d\n", int64(float64(hashes)/time))
}

// RandomNumber f
func RandomNumber(seed uint64) uint64 {
	seed ^= seed << 21
	seed ^= seed >> 31
	seed ^= seed << 4
	return seed
}

// RandomString x
func RandomString(len int, seed uint64) (uint64, []byte) {
	str := make([]byte, len)
	for i := 0; i < len; i++ {
		seed = RandomNumber(seed)
		str[i] = characterSet[seed%62]
	}
	return seed, str
}

// Hash func
func Hash(input []byte, challengeLength int) bool {
	hash := sha256.Sum256(input)
	for i := 0; i < challengeLength; i++ {
		if hash[i] > 0 {
			return false
		}
	}
	return true
}
