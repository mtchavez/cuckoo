package cuckoo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_New_withDefaults(t *testing.T) {
	filter := New()
	assert.IsType(t, filter.hasher, defaultHasher)
	assert.Equal(t, filter.kicks, defaultKicks)
	assert.Equal(t, filter.fingerprintLength, defaultFingerprintLength)
	assert.Equal(t, filter.bucketEntries, defaultBucketEntries)
	assert.Equal(t, filter.bucketTotal, defaultBucketTotal)
}

func TestFilter_New_withConfigOptions(t *testing.T) {
	kicks := uint(42)
	kicksOption := Kicks(kicks)

	entries := uint(42)
	entriesOption := BucketEntries(entries)

	buckets := uint(42)
	bucketsOption := BucketTotal(buckets)

	filter := New(
		kicksOption,
		entriesOption,
		bucketsOption,
	)

	assert.Equal(t, filter.hasher, defaultHasher)
	assert.Equal(t, filter.kicks, kicks)
	assert.Equal(t, filter.fingerprintLength, defaultFingerprintLength)
	assert.Equal(t, filter.bucketEntries, entries)
	assert.Equal(t, filter.bucketTotal, buckets)
}

func TestItemCount(t *testing.T) {
	filter := New()
	assert.Equal(t, filter.ItemCount(), uint(0))
	filter.count++
	assert.Equal(t, filter.ItemCount(), uint(1))
}
