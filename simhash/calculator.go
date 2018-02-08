package simhash

import (
	"crypto/md5"
	"strings"
)

const hashLength = 128

type SimHash [16]byte

func Calculate(text string) SimHash {
	words := strings.Split(strings.Replace(text, "\n", " ", -1), " ")

	var tmpRes [hashLength]int
	for _, w := range words {
		curHash := md5.Sum([]byte(w))
		for i := 0; i < 16; i++ {
			var mask byte = 0x80 // 0b10000000
			for j := 0; j < 8; j++ {
				if curHash[i]&mask != 0 {
					tmpRes[i*8+j] += 1
				} else {
					tmpRes[i*8+j] -= 1
				}
				mask >>= 1

			}
		}
	}

	var res SimHash
	var curByte byte = 0
	for i := 0; i < hashLength; i++ {
		if tmpRes[i] > 0 {
			curByte |= 0x1
		}

		if (i+1)%8 == 0 {
			res[i/8] = curByte
			curByte = 0
		} else {
			curByte <<= 1
		}
	}

	return res
}

func Difference(x, y SimHash) float32 {
	var equal = 0
	for i, v := range x {
		var xor = v ^ y[i]
		var mask byte = 0x1
		for j := 0; j < 8; j++ {
			if xor & mask == 0 {
				equal++
			}
			mask <<= 1
		}
	}

	return 1 - float32(equal) / float32(hashLength)
}
