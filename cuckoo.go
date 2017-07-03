package cuckoo

import (
	"encoding/binary"
	"hash"

	farm "github.com/dgryski/go-farm"
)

const magicNumber uint64 = 0x5bd1e995

type Filter struct {
	buckets           []bucket
	bucketEntries     uint
	bucketTotal       uint
	capacity          uint
	count             uint
	fingerprintLength uint
	hasher            hash.Hash
	kicks             uint
}

func New(opts ...configOption) (filter *Filter) {
	filter = &Filter{}
	for _, option := range opts {
		option(filter)
	}
	filter.configureDefaults()
	filter.createBuckets()
	return
}

func (f *Filter) Insert(item []byte) bool {
	fp := newFingerprint(item, f.fingerprintLength, f.hasher)
	i1 := uint(farm.Hash64(item)) % f.capacity
	i2 := f.alternateIndex(fp, i1)
	if f.insert(fp, i1) || f.insert(fp, i2) {
		return true
	}
	return f.relocationInsert(fp, i2)
}

func (f *Filter) ItemCount() uint {
	return f.count
}

func (f *Filter) insert(fp fingerprint, idx uint) bool {
	if f.buckets[idx].insert(fp) {
		f.count++
		return true
	}
	return false
}

func (f *Filter) relocationInsert(fp fingerprint, i uint) bool {
	for k := uint(0); k < f.kicks; k++ {
		f.buckets[i].relocate(fp)
		i = f.alternateIndex(fp, i)
		if f.insert(fp, i) {
			return true
		}
	}
	return false
}

func (f *Filter) createBuckets() {
	buckets := make([]bucket, f.capacity, f.capacity)
	for i := range buckets {
		buckets[i] = make([]fingerprint, f.bucketEntries, f.bucketEntries)
	}
	f.buckets = buckets
}

func (f *Filter) alternateIndex(fp fingerprint, i uint) uint {
	bytes := make([]byte, 64, 64)
	for i, b := range fp {
		bytes[i] = b
	}

	hash := binary.LittleEndian.Uint64(bytes)
	return uint(uint64(i)^(hash*magicNumber)) % f.capacity
}
