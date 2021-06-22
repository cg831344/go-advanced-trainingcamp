package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/internal/router"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func App() *http.Server {
	r := gin.Default()
	v1 := r.Group("/v1")
	router.InitAddUser(v1)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	return srv

}
