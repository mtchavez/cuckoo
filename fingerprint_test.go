package cuckoo

import (
	"hash/fnv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testHasher = fnv.New64()

func Test_newFingerprint(t *testing.T) {
	item := []byte("cuckoo")
	length := uint(6)
	fingerprinted := newFingerprint(item, length, testHasher)
	assert.Equal(t, fingerprinted, fingerprint{0x38, 0x72, 0x64, 0x6c, 0xda, 0xf0})
}
