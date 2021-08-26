package util

import (
	"github.com/pkg/errors"
	"strconv"
)

// ShouldInt64ID 验证并转换int64字符串ID
func ShouldInt64ID(id string) (int64ID int64, err error) {
	if id == "" {
		err = errors.New("id required")
		return
	}

	int64ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "invalid int64 id")
		return
	}

	return
}
