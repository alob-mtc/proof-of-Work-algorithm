package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
)

var characterSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var hashes = []int{}

func main() {
	start := time.Now()
	prefix := []byte("Bola Is A Boy")
	numberOfGoRHasher := 4
	numberOfGoRGen := 1
	solutionChan := make(chan []byte)
	// spin up the process
	// for i := 0; i < 3; i++ {
		POWOnCore(prefix, numberOfGoRGen, numberOfGoRHasher, solutionChan)
	// }
	solution := <-solutionChan
	fmt.Println(string(solution))
	var totalHashesProcessed int
	for _, val := range hashes {
		totalHashesProcessed += val
	}
	end := time.Now()
	time := end.Sub(start).Seconds()
	fmt.Println("total time:", time)
	fmt.Println("No of hashes:", humanize.Comma(int64(totalHashesProcessed)))
	fmt.Printf("processed/sec: %s\n", humanize.Comma(int64(float64(totalHashesProcessed)/time)))
}

// POWOnCore x
func POWOnCore(prefix []byte, numberOfGan, numberOfHasher int, solutionChan chan []byte) {
	blockSize := 1024
	size := numberOfHasher * 2
	// the precess channels
	unprocessIndex := make(chan int, size)
	processIndex := make(chan int, size)

	// blocks
	blocks := make([][][]byte, size)
	offset := len(prefix)
	for idx := range blocks {
		unprocessIndex <- idx
		blocks[idx] = make([][]byte, blockSize)
		for i := 0; i < blockSize; i++ {
			blocks[idx][i] = make([]byte, 20)
			blocks[idx][i] = append(prefix, blocks[idx][i]...)
		}
	}

	// Random string generator
	for i := 0; i < numberOfGan; i++ {
		go func() {
			seed := uint64(time.Now().Local().UnixNano())
			for blockIndex := range unprocessIndex {
				for _, val := range blocks[blockIndex] {
					seed = RandomString(val, offset, seed)
				}
				processIndex <- blockIndex
			}
		}()
	}
	// spin up the hashers
	for i := 0; i < numberOfHasher; i++ {
		index := len(hashes)
		hashes = append(hashes, 0)
		go func(hashIndex int) {
			for blockIndex := range processIndex {
				hashes[hashIndex] += blockSize
				for _, val := range blocks[blockIndex] {
					if Hash(val, 27) {
						solutionChan <- val
						break
					}
				}
				// return the index to the unprocessIndex chan => a new set of [prefix+randomString]byte is generated
				unprocessIndex <- blockIndex
			}
		}(index)
	}
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
