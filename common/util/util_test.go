package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenSortedColumn(t *testing.T) {
	type T1 struct {
		Name string
		Age  int
	}

	t1 := T1{
		Name: "asdf",
		Age:  121,
	}
	bytes, err := GenBytesOrderByColumn(t1)
	if err != nil {
		t.Failed()
	}
	assert.Equal(t, string(bytes), "Age=121;Name=asdf;")
}
