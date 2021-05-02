package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var ErrSignalExit = errors.New("signal exit")

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	//app服务
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello world")
	})
	appSrv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	//启动app服务
	g.Go(func() error {
		log.Println("start serveApp")
		return appSrv.ListenAndServe()
	})
	//停止服务
	g.Go(func() error {
		<-ctx.Done()
		log.Println("stop serveApp")
		return appSrv.Shutdown(ctx)
	})

	//debug服务
	debugSrv := http.Server{
		Addr:    ":8081",
		Handler: http.DefaultServeMux,
	}
	//启动debug服务
	g.Go(func() error {
		log.Println("start serveDebug")
		return debugSrv.ListenAndServe()
	})
	//停止服务
	g.Go(func() error {
		<-ctx.Done()
		log.Println("stop serveDebug")
		return debugSrv.Shutdown(ctx)
	})
	//启动监听退出信号
	log.Println("start signal")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-quit:
				//去退出serveApp服务和serveDebug服务
				return ErrSignalExit
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, ErrSignalExit) {
		log.Fatalf("server exception:%v", err)
	} else {
		log.Println("server exit")
	}
}
