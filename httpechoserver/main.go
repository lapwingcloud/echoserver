package httpechoserver

import (
	"fmt"
	"net/http"
	"os"
)

type StartOption struct {
	Bind      string
	LogFormat string
}

func Start(option StartOption) {
	logger := newLogger(option.LogFormat)

	logger.Info(fmt.Sprintf("http server listening at %v", option.Bind))
	err := http.ListenAndServe(option.Bind, &httpServer{logger: logger})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(1)
	}
}
