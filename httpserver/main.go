package httpechoserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/lapwingcloud/echoserver/util"
)

type StartOption struct {
	Bind      string
	LogFormat string
}

func Start(option StartOption) {
	version := util.Version()
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("failed to retrieve hostname: %v", err)
	}
	logger := util.NewLogger(
		option.LogFormat,
		slog.String("version", version),
		slog.String("hostname", hostname),
	)

	s := http.Server{
		Addr: option.Bind,
		Handler: &httpServer{
			logger:   logger,
			version:  version,
			hostname: hostname,
		},
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		sig := <-ch
		logger.Info(fmt.Sprintf("got signal %v, shutting down http server", sig))
		err := s.Shutdown(context.Background())
		if err != nil {
			logger.Error(fmt.Sprintf("failed to shutdown http server: %v", err))
		}
		wg.Done()
	}()

	logger.Info(fmt.Sprintf("http server listening at %v", s.Addr))
	err = s.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		wg.Wait()
		logger.Info("http server has shut down")
	} else {
		logger.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(1)
	}
}
