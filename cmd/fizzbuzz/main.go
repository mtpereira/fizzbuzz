package main

import (
	"context"
	"fizzbuzz/endpoint"
	"fizzbuzz/middleware"
	"fizzbuzz/service"
	"fizzbuzz/transport"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-kit/kit/log"
)

func main() {
	httpListenAddress := flag.String("http", ":8080", "http listen address")
	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	s := service.New()
	e := endpoint.New(s)
	e.Single = middleware.Logging(log.With(logger, "method", "single"))(e.Single)
	fbv1 := transport.NewHTTPHandler(*e, logger)

	srv := &http.Server{Addr: *httpListenAddress, Handler: fbv1}

	// Handle SIGINT.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	(wg).Add(1)
	go func() {
		defer (wg).Done()

		go func() {
			logger.Log("msg", "http", "action", "start", "addr", *httpListenAddress)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("msg", "http", "action", "stop", "addr", *httpListenAddress)

		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()

	logger.Log("msg", "exit", "reason", <-errc)
	cancel()
	wg.Wait()
}
