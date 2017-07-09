package cuckoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newFingerprint(t *testing.T) {
	item := []byte("cuckoo")
	length := uint(6)
	fingerprinted := newFingerprint(item, length)
	assert.Equal(t, fingerprinted, fingerprint{0x72, 0x65, 0x0, 0x31, 0xc2, 0xe})
}
