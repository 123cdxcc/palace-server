package errors

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	fmt.Printf("%+v", f2())
}

func f2() error {
	return f1()
}

func f1() error {
	return New("错误")
}
