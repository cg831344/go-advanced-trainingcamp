package timeout

import (
	"fmt"
	context "golang.org/x/net/context"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()

	<-ctx.Done()
	fmt.Println(ctx.Err())

}
