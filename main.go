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
	random := make([]byte, 20)
	random = append(prefix, random...)
	for {
		hashes++
		seed, str = RandomString(random, len(prefix), seed)
		if Hash(random, 3) {
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
func RandomString(str []byte, offset int, seed uint64) (uint64, []byte) {
	for i := offset; i < len(str); i++ {
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
