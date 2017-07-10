package cuckoo

import "fmt"

func ExampleFilter_Save() {
	filter := New()
	item := []byte("Largo")
	filter.InsertUnique(item)
	filter.Save("./tmp/example_save.gob")

	loadedFilter, _ := Load("./tmp/example_save.gob")
	fmt.Printf("Loaded filter has same item? %v\n\n", loadedFilter.Lookup(item))
	fmt.Printf("[Old]\nBucketEntries: %v\nBucketTotal: %v\nFingerprint: %v\nKicks: %v\n\n", filter.bucketEntries, filter.bucketTotal, filter.fingerprintLength, filter.kicks)
	fmt.Printf("[Loaded]\nBucketEntries: %v\nBucketTotal: %v\nFingerprint: %v\nKicks: %v\n", loadedFilter.bucketEntries, loadedFilter.bucketTotal, loadedFilter.fingerprintLength, loadedFilter.kicks)
	// Output:
	// Loaded filter has same item? true
	//
	// [Old]
	// BucketEntries: 24
	// BucketTotal: 4194304
	// Fingerprint: 6
	// Kicks: 500
	//
	// [Loaded]
	// BucketEntries: 24
	// BucketTotal: 4194304
	// Fingerprint: 6
	// Kicks: 500
}
