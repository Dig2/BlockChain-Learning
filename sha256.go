package main

import (
	"encoding/binary"
	"fmt"
)

func prepare(s []byte) [][]byte {
	// append a bytes started with bit '1'
	tmp := append(s, 0x80)
	l := len(s)
	// zero-padding
	if l % 64 < 56 {
		tmp = append(tmp, make([]byte, 55 - l % 64)...)
	} else {
		tmp = append(tmp, make([]byte, 121 - l % 64)...)
	}
	// raw data length
	tmp2 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp2, uint64(l<<3))
	tmp = append(tmp, tmp2...)
	// chunk
	var prepared [][]byte
	for i := 0; i < len(tmp) / 64; i++ {
		prepared = append(prepared, tmp[i * 64 : i * 64 + 64])
	}
	return prepared
}

func sha256(chunks [][]byte) []byte {
	for _, chunk := range chunks {
		//fmt.Println(chunk)
		w := make([]uint32, 64)
		for i := 0; i < 16; i++ {
			w[i] = binary.BigEndian.Uint32(chunk[i * 4: i * 4 + 4])
		}
		for i := 16; i < 64; i++ {
			s0 := (rightRotate(w[i-15], 7)) ^ (rightRotate(w[i-15], 18)) ^ (w[i-15] >> 3)
			s1 := (rightRotate(w[i-2], 17)) ^ (rightRotate(w[i-2], 19)) ^ (w[i-2] >> 10)
			w[i] = w[i-16] + s0 + w[i-7] + s1
		}
		fmt.Println(w[16])
	}
	// TODO
	return []byte("aa")
}

func rightRotate(n uint32, d uint) uint32 {
	return (n >> d) | (n << (32 - d))
}

func main()  {

	initHash := [...]uint32{0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19}
	initConst := [...]uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5, 0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174 ,0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da ,0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967 ,0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85 ,0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070 ,0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3 ,0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

	fmt.Println(initConst)
	fmt.Println(initHash)

	s := "hello world"
	prepared := prepare([]byte(s))
	res := sha256(prepared)
	fmt.Println(res)
}
