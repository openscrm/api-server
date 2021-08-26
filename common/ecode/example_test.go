package ecode_test

import (
	"fmt"
	"github.com/pkg/errors"
	ecode2 "openscrm/common/ecode"
)

func Example_ecode_Message() {
	_ = ecode2.OK.Message()
}

func Example_ecode_Code() {
	_ = ecode2.OK.Code()
}

func Example_ecode_Error() {
	_ = ecode2.OK.Error()
}

func ExampleCause() {
	err := errors.WithStack(ecode2.InternalError)
	ecode2.Cause(err)
}

func ExampleInt() {
	err := ecode2.Int(500)
	fmt.Println(err)
	// Output:
	// 500
}

func ExampleString() {
	ecode2.String("500")
}

// ExampleStack package error with stack.
func Example() {
	err := errors.New("dao error")
	errors.Wrap(err, "some message")
	// package ecode with stack.
	errCode := ecode2.InternalError
	err = errors.Wrap(errCode, "some message")

	//get ecode from package error
	code := errors.Cause(err).(ecode2.Codes)
	fmt.Printf("%d: %s\n", code.Code(), code.Message())
}
