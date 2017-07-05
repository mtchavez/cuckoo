package cuckoo

import (
	"bufio"
	"os"
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

func TestFilter_New_withLowCapacity(t *testing.T) {
	entries := uint(1 << 18)
	entriesOption := BucketEntries(entries)

	buckets := uint(4)
	bucketsOption := BucketTotal(buckets)

	filter := New(
		entriesOption,
		bucketsOption,
	)

	assert.Equal(t, filter.capacity, uint(1))
	assert.Equal(t, filter.bucketEntries, entries)
	assert.Equal(t, filter.bucketTotal, buckets)
}

func TestInsert(t *testing.T) {
	filter := New()
	fd, err := os.Open("/usr/share/dict/words")
	defer fd.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	scanner := bufio.NewScanner(fd)
	var wordCount uint
	var totalWords uint
	var values [][]byte
	for scanner.Scan() {
		word := []byte(scanner.Text())
		totalWords++

		if filter.Insert(word) {
			wordCount++
		}
		values = append(values, word)
	}

	assert.Equal(t, int(filter.ItemCount()), int(totalWords))
}

func TestInsert_withRelocations(t *testing.T) {
	filter := New(
		BucketTotal(250000),
		BucketEntries(6),
		FingerprintLength(1),
	)

	fd, err := os.Open("/usr/share/dict/words")
	defer fd.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	scanner := bufio.NewScanner(fd)
	var wordCount uint
	var totalWords uint
	var values [][]byte
	for scanner.Scan() {
		word := []byte(scanner.Text())
		totalWords++

		if filter.Insert(word) {
			wordCount++
		}
		values = append(values, word)
	}
	inserted := int(filter.ItemCount())
	total := int(totalWords)
	miss := 1 - (float64(inserted) / float64(total))
	assert.Equal(t, miss <= 0.065, true)
}

func TestInsertUnique(t *testing.T) {
	filter := New()
	item := []byte("an-item")
	filter.InsertUnique(item)
	assert.Equal(t, filter.ItemCount(), uint(1))
	filter.InsertUnique(item)
	assert.Equal(t, filter.ItemCount(), uint(1))
}

func TestInsertUnique_withNewItem(t *testing.T) {
	filter := New()
	filter.InsertUnique([]byte("an-item"))
	assert.Equal(t, filter.ItemCount(), uint(1))
	filter.InsertUnique([]byte("another-item"))
	assert.Equal(t, filter.ItemCount(), uint(2))
}

func TestItemCount(t *testing.T) {
	filter := New()
	assert.Equal(t, filter.ItemCount(), uint(0))
	filter.count++
	assert.Equal(t, filter.ItemCount(), uint(1))
}

func TestLookup(t *testing.T) {
	filter := New()
	fd, err := os.Open("/usr/share/dict/words")
	defer fd.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	scanner := bufio.NewScanner(fd)
	var wordCount uint
	var totalWords uint
	var values [][]byte
	for scanner.Scan() {
		word := []byte(scanner.Text())
		totalWords++

		if filter.Insert(word) {
			wordCount++
		}
		values = append(values, word)
	}

	var found = 0
	for _, word := range values {
		if !filter.Lookup(word) {
			found++
			t.Errorf("Expected to find %+v in filter", word)
		}
	}
	total := int(totalWords)
	miss := float64(found) / float64(total)
	assert.Equal(t, miss, float64(0))
}

func TestLookup_notFound(t *testing.T) {
	filter := New()
	item := []byte("an-item")
	filter.InsertUnique(item)
	found := filter.Lookup([]byte("not-found"))
	assert.Equal(t, found, false)
}

func TestDelete(t *testing.T) {
	filter := New()
	fd, err := os.Open("/usr/share/dict/words")
	defer fd.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	scanner := bufio.NewScanner(fd)
	var wordCount uint
	var totalWords uint
	var values [][]byte
	for scanner.Scan() {
		word := []byte(scanner.Text())
		totalWords++

		if filter.Insert(word) {
			wordCount++
		}
		values = append(values, word)
	}

	for _, word := range values {
		if !filter.Delete(word) {
			t.Errorf("Expected to delete %+v in filter", word)
		}
	}
	assert.Equal(t, filter.ItemCount(), uint(0))
}

func TestDelete_notFound(t *testing.T) {
	filter := New()
	item := []byte("an-item")
	filter.InsertUnique(item)
	deleted := filter.Delete([]byte("not-found"))
	assert.Equal(t, deleted, false)
}
