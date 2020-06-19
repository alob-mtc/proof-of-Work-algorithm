package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var characterSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func main() {
	start := time.Now()
	prefix := []byte("Bola is a boy")
	numberOfGoR := 4
	hashes := make([]int, numberOfGoR)
	solutionChan := make(chan []byte, 1)

	for i := 0; i < numberOfGoR; i++ {
		go func(index int) {
			offset := len(prefix)
			random := make([]byte, 20)
			random = append(prefix, random...)
			seed := uint64(index + 1)
			for {
				hashes[index]++
				seed = RandomString(random, offset, seed)
				if Hash(random, 27) {
					solutionChan <- random
					break
				}
			}
		}(i)
	}

	solution := <-solutionChan
	fmt.Println(string(solution))
	var totalHashesProcessed int
	for _, val := range hashes {
		totalHashesProcessed += val
	}
	end := time.Now()
	time := end.Sub(start).Seconds()
	fmt.Println("total time:", time)
	fmt.Println("No of hashes:", totalHashesProcessed)
	fmt.Printf("processed/sec: %d\n", int64(float64(totalHashesProcessed)/time))
}

// RandomNumber f
func RandomNumber(seed uint64) uint64 {
	seed ^= seed << 21
	seed ^= seed >> 31
	seed ^= seed << 4
	return seed
}

// RandomString x
func RandomString(str []byte, offset int, seed uint64) uint64 {
	for i := offset; i < len(str); i++ {
		seed = RandomNumber(seed)
		str[i] = characterSet[seed%62]
	}
	return seed
}

// Hash func
// ? i don't know how this fully works => checking bit-wise
func Hash(input []byte, bits int) bool {
	hash := sha256.Sum256(input)
	nbytes := bits / 8
	nbits := bits % 8
	var idx int //this init it to 0
	for ; idx < nbytes; idx++ {
		if hash[idx] > 0 {
			return false
		}
	}
	return (hash[idx] >> (8 - nbits)) == 0
}
