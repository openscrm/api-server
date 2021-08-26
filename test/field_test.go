package test

import (
	"github.com/stretchr/testify/assert"
	"openscrm/app/constants"
	"testing"
)

func TestTimeField(t *testing.T) {
	var val constants.TimeField
	err := val.UnmarshalJSON([]byte(`"08:01:01"`))
	assert.Nil(t, err)
	t.Log(val)
	t.Log(val.Seconds())
}
