package main

import (
	"errors"
	"log"
)

var ErrParams = errors.New("params error")

type WrapError struct {
	msg string
	err error
}

func (w *WrapError) Error() string {
	return w.msg
}

// Unwrap 实现这个方法可以通过Is判断错误
func (w *WrapError) Unwrap() error {
	return w.err
}

func main() {
	err := &WrapError{
		msg: "wrap error msg",
		err: ErrParams,
	}
	log.Println(errors.Is(err, ErrParams))

	var we *WrapError
	// As 用于断言错误
	if errors.As(err, &we) {
		log.Println(we.msg)
	}

}
