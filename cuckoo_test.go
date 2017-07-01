package cuckoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_New_withDefaults(t *testing.T) {
	filter := New()
	assert.IsType(t, filter.hasher, defaultHasher)
	assert.IsType(t, filter.kicks, defaultKicks)
	assert.IsType(t, filter.fingerprintLength, defaultFingerprintLength)
	assert.IsType(t, filter.bucketEntries, defaultBucketEntries)
	assert.IsType(t, filter.bucketTotal, defaultBucketTotal)
}

func TestFilter_New_withConfigOptions(t *testing.T) {
	kicks := uint(42)
	kicksOption := Kicks(kicks)

	entries := uint(42)
	entriesOption := BucketEntries(entries)

	buckets := uint(42)
	bucketsOption := BucketEntries(buckets)

	filter := New(
		kicksOption,
		entriesOption,
		bucketsOption,
	)

	assert.IsType(t, filter.hasher, defaultHasher)
	assert.IsType(t, filter.kicks, kicks)
	assert.IsType(t, filter.fingerprintLength, defaultFingerprintLength)
	assert.IsType(t, filter.bucketEntries, entries)
	assert.IsType(t, filter.bucketTotal, buckets)
}
