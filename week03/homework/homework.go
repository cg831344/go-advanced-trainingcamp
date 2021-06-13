package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 监控信号量，返回一个CancelContext
func NotifySingle() context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-c:
			cancel()
		}
	}()
	return ctx
}

func appHttpServer() *http.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, HandlerFunc!")
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(f))

	srv := &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}

	return srv

}

func debugHttpServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "okay")
		},
	))

	srv := &http.Server{
		Addr:    ":9001",
		Handler: mux,
	}
	return srv
}

func shutDownServer(ctx context.Context, srvs ...*http.Server) {
	for _, srv := range srvs {
		_ = srv.Shutdown(ctx)

	}
}

func main() {
	//注册信号量
	ctx := NotifySingle()

	appHttpServer := appHttpServer()
	debugHttpServer := debugHttpServer()
	var g errgroup.Group

	g.Go(func() error {

		if err := appHttpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "appHttpServer listen fail")
		}
		return nil
	})
	g.Go(func() error {

		if err := debugHttpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "debugHttpServer listen fail")
		}
		return nil

	})

	// 监听到退出信号后，关闭所有http src
	go func() {
		select {
		case <-ctx.Done():
			log.Println("收到关闭信号")
			shutDownServer(ctx, appHttpServer, debugHttpServer)
		}
	}()

	err := g.Wait()

	if err != nil {
		if errors.As(err, &http.ErrServerClosed) {
			log.Printf("正常关闭")
		} else {
			log.Printf("启动异常：%s", err)
			shutDownServer(ctx, appHttpServer, debugHttpServer)
		}

	}

}
