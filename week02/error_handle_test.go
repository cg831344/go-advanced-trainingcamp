package week02

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"testing"
)

type MyError struct {

}

func (receiver MyError) Error() string {
	return "这是我的错"

}


func openFile() error {
	_,err:=os.Open("/tmp/a.txt")
	if err != nil {
		return errors.Wrap(err,"打开失败")
	}
	return nil
}

func returnError() error  {
	return MyError{}
}

func TestName(t *testing.T) {
	err := openFile()
	err = errors.Wrap(err,"能打开吗")
	fmt.Printf("%T,%v\n",errors.Cause(err),errors.Cause(err))
	fmt.Printf("%+v",err)

}

func TestError(t *testing.T) {
	err := returnError()
	if ok := errors.As(err,&MyError{});ok {
		fmt.Printf("ok\n", )
	}

}
