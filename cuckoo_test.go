package cuckoo

import (
	"testing"
)

func Test(t *testing.T) {
	if true != true {
		t.Errorf("Boom")
	}
}
