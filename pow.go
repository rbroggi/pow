package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
)

// PoW is a simple proof of work function that takes a byte slice and a difficulty
// and returns a nonce that satisfies the difficulty.
// The difficulty is the number of leading zeros bits in the hash of the input and nonce.
func PoW(input []byte, difficulty int) uint64 {
	for nonce := uint64(0); ; nonce++ {
		hash := sha256.Sum256(append(Uint64ToBytes(nonce), input...))
		if checkDifficulty(hash[:], difficulty) {
			return nonce
		}
		nonce++
	}
}

func PoWWithLikelihood(input []byte, difficulty int, likelihood float64) (uint64, error) {
	maxAttempts := math.Ceil(math.Log(1-likelihood) / math.Log(1.-(1.0/math.Pow(2, float64(difficulty)))))
	for nonce := uint64(0); nonce < uint64(maxAttempts); nonce++ {
		hash := sha256.Sum256(append(Uint64ToBytes(nonce), input...))
		if checkDifficulty(hash[:], difficulty) {
			return nonce, nil
		}
	}
	return 0, fmt.Errorf("failed to find nonce with likelihood %f", likelihood)
}

func VerifyPoW(input []byte, nonce uint64, difficulty int) bool {
	hash := sha256.Sum256(append(Uint64ToBytes(nonce), input...))
	return checkDifficulty(hash[:], difficulty)
}

func Uint64ToBytes(n uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, n)
	return bytes
}

// checkDifficulty checks if the hash has the required number of leading zeros bits.
func checkDifficulty(hash []byte, difficulty int) bool {
	// check if the first difficulty bits are zeros
	for i := 0; i < difficulty; i++ {
		// hash[i/8] gets the byte that contains the i-th bit
		// 1<<(7-uint(i%8)) creates a byte with the i-th bit set
		// the two are ANDed to check if the i-th bit is set
		if hash[i/8]&(1<<(7-uint(i%8))) != 0 {
			return false
		}
	}
	return true
}
