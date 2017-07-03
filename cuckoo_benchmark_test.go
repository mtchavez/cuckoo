package cuckoo

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func BenchmarkCuckooNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkInsert(b *testing.B) {
	filter := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Insert([]byte(fmt.Sprintf("item-%d", i%50000)))
	}
}

func BenchmarkInsertUnique(b *testing.B) {
	filter := New()
	fd, _ := os.Open("/usr/share/dict/words")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var wordCount int
	var totalWords int
	var values [][]byte
	for scanner.Scan() {
		word := []byte(scanner.Text())
		totalWords++

		if filter.Insert(word) {
			wordCount++
		}
		values = append(values, word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.InsertUnique(values[i%totalWords])
	}
}

func BenchmarkLookup(b *testing.B) {
	filter := New()
	fd, _ := os.Open("/usr/share/dict/words")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	var wordCount int
	var totalWords int
	var values [][]byte
	for scanner.Scan() {
		word := []byte(scanner.Text())
		totalWords++

		if filter.Insert(word) {
			wordCount++
		}
		values = append(values, word)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Lookup(values[i%totalWords])
	}
}
