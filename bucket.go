package cuckoo

import (
	"math/rand"
)

type bucket []fingerprint

func (b bucket) insert(fp fingerprint) bool {
	for i, fprint := range b {
		if fprint == nil {
			b[i] = fp
			return true
		}
	}
	return false
}

func (b bucket) relocate(fp fingerprint) fingerprint {
	i := rand.Intn(len(b))
	b[i], fp = fp, b[i]

	return fp
}
