package test

import (
	"github.com/stretchr/testify/assert"
	"openscrm/app/constants"
	"openscrm/common/storage"
	"testing"
)

func TestOSSStorage_SignURL(t *testing.T) {
	signed, err := storage.FileStorage.SignURL("test/test.pdf", constants.HTTPPut, 86400)
	assert.Nil(t, err)
	t.Log(signed)
}
