package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	g := errgroup.Group{}

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// context withCancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// define mutilple handlers
	sm := http.NewServeMux()
	sm.HandleFunc("/hello", helloHandler)

	// 生产服务
	httpServerProduct := http.Server{
		Handler: sm,
		Addr:    ":8080",
	}

	// 调试服务
	httpServerDebug := http.Server{
		Handler: sm,
		Addr:    ":8090",
	}

	g.Go(func() error {
		logrus.Infof("Product Env Listen on %v\n", httpServerProduct.Addr)
		return httpServerProduct.ListenAndServe()
	})

	g.Go(func() error {
		logrus.Infof("Debug Env Listen on %v\n", httpServerDebug.Addr)
		return httpServerDebug.ListenAndServe()
	})

	// 信号捕获和关闭httpServer
	g.Go(func() error {
		var msg string
		select {
		case s := <-signalCh:
			msg = fmt.Sprintf("Got signal: %s", s)
		case <-ctx.Done():
			msg = fmt.Sprintf("ctx err: %v", ctx.Err())
		}

		if err := httpServerProduct.Shutdown(ctx); err != nil {
			return err
		}
		if err := httpServerDebug.Shutdown(ctx); err != nil {
			return err
		}
		return errors.New(msg)
	})

	// 阻塞 Wait for
	if err := g.Wait(); err != nil {
		logrus.Warnf("error: %s", err)
	}
}
