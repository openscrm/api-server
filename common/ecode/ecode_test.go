package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	defer func() {
		errStr := recover()
		if errStr != "ecode: 2810001 already exist" {
			t.Logf("New duplicate ecode should cause panic")
			t.FailNow()
		}
	}()
	var _ error = New(2810001)
	var _ error = New(2810002)
	var _ error = New(2810001)
}

func TestErrMessage(t *testing.T) {
	e1 := New(2810003)
	assert.Equal(t, 2810003, e1.Code())
	assert.Equal(t, "2810003", e1.Message())
	RegisterMessages(map[int]Message{2810003: Message{
		Msg:    "testErr",
		Detail: "测试错误",
	}})
	assert.Equal(t, "testErr", e1.LocalizedMessage(En))
	assert.Equal(t, "测试错误", e1.LocalizedMessage(Zh))
}

func TestCause(t *testing.T) {
	e1 := New(2810004)
	err := errors.Wrap(e1, "wrap error")
	e2 := Cause(err)
	assert.Error(t, e1, e2)
}

func TestMessage(t *testing.T) {
	for code, msg := range _messages {
		fmt.Printf("%d	%s	%s\n", code, msg.Msg, msg.Detail)
	}
}
