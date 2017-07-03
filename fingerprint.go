package cuckoo

import (
	"hash"
)

type fingerprint []byte

func newFingerprint(item []byte, length uint, hasher hash.Hash) fingerprint {
	hasher.Reset()
	hasher.Write(item)
	hashedFingerprint := hasher.Sum(nil)

	fingerprinted := make(fingerprint, length, length)
	for i := uint(0); i < length; i++ {
		fingerprinted[i] = hashedFingerprint[i]
	}

	if fingerprinted == nil {
		fingerprinted[0] += 7
	}
	return fingerprinted
}
