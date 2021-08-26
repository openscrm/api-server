package util

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/pkg/errors"
	"openscrm/app/constants"
)

// GetInt64FromSession 从指定Session中获取int64字段
func GetInt64FromSession(sess sessions.Session, fieldName constants.SessionField) (val int64, err error) {
	id, ok := sess.Get(string(fieldName)).(string)
	if !ok {
		err = errors.New("invalid " + string(fieldName))
		return
	}

	val, err = ShouldInt64ID(id)
	if err != nil {
		err = errors.Wrap(err, "ShouldInt64ID failed")
		return
	}

	return
}

// GetInt64IDFromSession 从指定Session中获取int64ID
func GetInt64IDFromSession(sess sessions.Session, fieldName constants.SessionField) (id string, err error) {
	val, err := GetInt64FromSession(sess, fieldName)
	if err != nil {
		err = errors.Wrap(err, "GetInt64FromSession failed")
		return
	}

	id = fmt.Sprintf("%d", val)

	return
}
