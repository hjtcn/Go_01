package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pkg/mod/golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c/errgroup"
	"syscall"
	"time"
)

/**
作业命题：
	基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/

const server_port_one = ":8081"
const server_port_two = ":8082"
const server_port_three = ":8083"

type httpServer struct {
	http	 *http.Server
	shutdown chan struct{}
}

func (h *httpServer) startHttpServer(ctx context.Context) (err error) {
	go func() {
		for {
			select {
			case <-h.shutdown:
				fmt.Println("即将关闭服务", h.http.Addr, "...")

				h.http.SetKeepAlivesEnabled(false)
				time.Sleep(2 * time.Second)

				if h.http.Shutdown(ctx) != nil {
					fmt.Println("已关闭服务", h.http.Addr)
					return
				}

				fmt.Println("服务关闭失败", h.http.Addr)

			default:
			}
		}
	}()

	return h.http.ListenAndServe()
}

func graceFullShutdown(quit <-chan os.Signal, Shutdown chan<- struct{}) (err error) {
	fmt.Println("收到关闭信号")
	<-quit

	if Shutdown == nil {
		panic("gorpc.Server: server must be started before stopping it")
	}

	//收到关闭信号，关闭chan
	close(Shutdown)
	return
}

func main() {
	shutdown := make(chan struct{})

	server_one := &httpServer{
		http:     &http.Server{
			Addr:              server_port_one,
		},
		shutdown: shutdown,
	}

	server_two := &httpServer{
		http:     &http.Server{
			Addr:              server_port_two,
		},
		shutdown: shutdown,
	}

	server_three := &httpServer{
		http:     &http.Server{
			Addr:              server_port_three,
		},
		shutdown: shutdown,
	}

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return server_one.startHttpServer(ctx)
	})

	g.Go(func() error {
		return server_two.startHttpServer(ctx)
	})

	g.Go(func() error {
		return server_three.startHttpServer(ctx)
	})

	g.Go(func() error {
		return graceFullShutdown(quit, shutdown)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("http服务异常，异常信息: ", err)
	}

	<-shutdown
	fmt.Println("all server is stop!!!")
}

/**
害，自己写还是不会，参照github名称为winkyi写的，思路大致明白了
 */
