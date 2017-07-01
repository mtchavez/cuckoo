package cuckoo

import (
	"hash"
)

type Filter struct {
	bucketEntries     uint
	bucketTotal       uint
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
	return
}

func (f *Filter) ItemCount() uint {
	return f.count
}
