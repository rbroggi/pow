package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestPoW(t *testing.T) {
	tests := map[string]struct {
		input      []byte
		difficulty int
	}{
		"test 1": {
			input:      randomBytes(t, 32),
			difficulty: 1,
		},
		"test 10": {
			input:      randomBytes(t, 32),
			difficulty: 10,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// print bit representation of the hash
			// fmt.Printf("%08b\n", sha256.Sum256(append(tt.input, Uint64ToBytes(tt.want)...)))
			got := PoW(tt.input, tt.difficulty)
			hash := sha256.Sum256(append(Uint64ToBytes(got), tt.input...))
			var binaryHashStr string
			for i := 0; i < len(hash); i++ {
				binaryHashStr += fmt.Sprintf("%08b", hash[i])
			}
			// assert that the hash has the required number of leading zeros bits
			for i := 0; i < tt.difficulty; i++ {
				if binaryHashStr[i] != '0' {
					t.Fatalf("Index %d should be 0", i)
				}
			}
			// print nonce
			fmt.Printf("binaryHashStr: %s\n", binaryHashStr)
			fmt.Printf("nonce: %d\n", got)
			if VerifyPoW(tt.input, got, tt.difficulty) != true {
				t.Fatalf("VerifyPoW() = false, want true")
			}
		})
	}
}

func TestPoWWithDifficulty(t *testing.T) {
	tests := map[string]struct {
		input       []byte
		difficulty  int
		Likelihood  float64
		expectedErr bool
	}{
		"test fail to find (can fail at times)": {
			input:       randomBytes(t, 32),
			difficulty:  20,
			Likelihood:  0.00001,
			expectedErr: true,
		},
		"passes - high likelihood": {
			input:       randomBytes(t, 32),
			difficulty:  5,
			Likelihood:  0.99,
			expectedErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// print bit representation of the hash
			// fmt.Printf("%08b\n", sha256.Sum256(append(tt.input, Uint64ToBytes(tt.want)...)))
			got, err := PoWWithLikelihood(tt.input, tt.difficulty, tt.Likelihood)
			if tt.expectedErr {
				if err == nil {
					t.Fatalf("PoWWithLikelihood() = %d, want error", got)
				}
				return
			}
			hash := sha256.Sum256(append(Uint64ToBytes(got), tt.input...))
			var binaryHashStr string
			for i := 0; i < len(hash); i++ {
				binaryHashStr += fmt.Sprintf("%08b", hash[i])
			}
			// assert that the hash has the required number of leading zeros bits
			for i := 0; i < tt.difficulty; i++ {
				if binaryHashStr[i] != '0' {
					t.Fatalf("Index %d should be 0", i)
				}
			}
			// print nonce
			fmt.Printf("binaryHashStr: %s\n", binaryHashStr)
			fmt.Printf("nonce: %d\n", got)
		})
	}
}

func randomBytes(t *testing.T, size int) []byte {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func randomBytesPanic(size int) []byte {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return bytes
}

func BenchmarkSHA256Hash(b *testing.B) {
	data := []byte("your data here")
	for i := 0; i < b.N; i++ {
		hash := sha256.New()
		hash.Write(data)
		hash.Sum(nil)
	}
}

func BenchmarkSHA256OpsPerSecond(b *testing.B) {
	data := []byte("sample payload")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hash := sha256.Sum256(data)
		_ = hash
	}
}

func BenchmarkPoW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := randomBytesPanic(32)
		PoW(data, 20)
	}
}
