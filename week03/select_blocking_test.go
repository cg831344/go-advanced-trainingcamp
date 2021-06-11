package week03

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestBlocking(t *testing.T) {
	ch := make(chan int, 1)
	f := func(ch chan int) {
		time.Sleep(time.Second * 5)
		ch <- 1
	}
	go f(ch)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("失败超时")
	case res := <-ch:
		fmt.Println("ok", res)
	}
}
