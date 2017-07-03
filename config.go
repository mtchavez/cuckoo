package cuckoo

import (
	"hash"
	"hash/fnv"
)

type configOption func(*Filter)

// Cuckoo Filter notation
// e target false positive rate
// f fingerprint length in bits
// α load factor (0 ≤ α ≤ 1)
// b number of entries per bucket
// m number of buckets
// n number of items
// C average bits per item
const (
	// Entries per bucket (b)
	defaultBucketEntries uint = 24
	// Bucket total (m) defaults to approx. 4 million
	defaultBucketTotal uint = 1 << 22
	// Length of fingreprint (f) set to log(n/b) ~6 bits
	defaultFingerprintLength uint = 6
	// Default attempts to find empty slot on insert
	defaultKicks uint = 500
)

var (
	defaultHasher = fnv.New64()
)

func BucketEntries(entries uint) configOption {
	return func(f *Filter) {
		f.bucketEntries = entries
	}
}

func BucketTotal(total uint) configOption {
	return func(f *Filter) {
		f.bucketTotal = total
	}
}

func FingerprintLength(length uint) configOption {
	return func(f *Filter) {
		f.fingerprintLength = length
	}
}

func Hasher(hasher hash.Hash) configOption {
	return func(f *Filter) {
		f.hasher = hasher
	}
}

func Kicks(kicks uint) configOption {
	return func(f *Filter) {
		f.kicks = kicks
	}
}

func capacity() configOption {
	return func(f *Filter) {
		f.capacity = nextPowerOf2(uint64(f.bucketTotal)) / f.bucketEntries
		if f.capacity <= 0 {
			f.capacity = 1
		}
	}
}

func (f *Filter) configureDefaults() {
	if f.bucketTotal <= 0 {
		BucketTotal(defaultBucketTotal)(f)
	}

	if f.bucketEntries <= 0 {
		BucketEntries(defaultBucketEntries)(f)
	}

	if f.fingerprintLength <= 0 {
		FingerprintLength(defaultFingerprintLength)(f)
	}

	if f.hasher == nil {
		Hasher(defaultHasher)(f)
	}

	if f.kicks <= 0 {
		Kicks(defaultKicks)(f)
	}

	if f.capacity < 1 {
		capacity()(f)
	}
}

func nextPowerOf2(n uint64) uint {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return uint(n)
}
