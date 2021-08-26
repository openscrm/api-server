package test

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInheritStructCopy(t *testing.T) {
	type A struct {
		Name string
	}

	type B struct {
		A
	}

	type B2 struct {
		A
	}

	type C struct {
		Name string
	}

	b := &B{
		A{Name: "b"},
	}

	b2 := &B2{
		A{Name: "b2"},
	}

	c := &C{
		Name: "c",
	}

	err := copier.Copy(b, c)
	assert.Nil(t, err)
	assert.Equal(t, "c", b.Name)

	err = copier.Copy(b, b2)
	assert.Nil(t, err)
	assert.Equal(t, "b2", b.Name)

	err = copier.Copy(c, b)
	assert.Nil(t, err)
	assert.Equal(t, "b", c.Name)
}
