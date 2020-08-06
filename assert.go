package assert

import (
	"errors"
	"fmt"
)

type (
	Bang interface {
		Panic()
		Error() error
	}

	bang struct {
		condFn  func() bool
		msgFmt  string
		msgArgs []interface{}
	}
)

func (this *bang) Panic() {
	if err := this.Error(); err != nil {
		panic(err)
	}
}

func (this *bang) Error() error {
	if this.condFn() {
		return nil
	}
	return errors.New(fmt.Sprintf(this.msgFmt, this.msgArgs...))
}

func makeBangFn(condFn func() bool, msgFmt string, msgArgs ...interface{}) Bang {
	return &bang{
		condFn:  condFn,
		msgFmt:  msgFmt,
		msgArgs: msgArgs,
	}
}

func makeBang(cond bool, msgFmt string, msgArgs ...interface{}) Bang {
	return makeBangFn(
		func() bool {
			return cond
		},
		msgFmt,
		msgArgs...)
}

func Must(cond bool, msg string) Bang {
	return Mustf(cond, msg)
}

func Mustf(cond bool, format string, args ...interface{}) Bang {
	return makeBang(cond, format, args...)
}
