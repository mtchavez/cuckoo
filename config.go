package cuckoo

import (
	"hash"
	"hash/fnv"
)

type configOption func(*Filter)

const (
	defaultBucketEntries     uint = 8
	defaultBucketTotal       uint = 2 ^ 30
	defaultFingerprintLength uint = 6
	defaultKicks             uint = 500
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

func (f *Filter) configureDefaults() {
	if f.bucketEntries <= 0 {
		BucketEntries(defaultBucketEntries)(f)
	}

	if f.bucketTotal <= 0 {
		BucketTotal(defaultBucketTotal)(f)
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
}
