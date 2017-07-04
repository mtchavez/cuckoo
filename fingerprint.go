package cuckoo

import (
	"hash"
	"sync"
)

type fingerprint []byte

var hashSync sync.Mutex

func newFingerprint(item []byte, length uint, hasher hash.Hash) fingerprint {
	hashSync.Lock()
	defer hashSync.Unlock()
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
