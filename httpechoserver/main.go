package httpechoserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lapwingcloud/echoserver/util"
)

type StartOption struct {
	Bind      string
	LogFormat string
}

func Start(option StartOption) {
	logger := util.NewLogger(option.LogFormat)

	logger.Info(fmt.Sprintf("http server listening at %v", option.Bind))
	err := http.ListenAndServe(option.Bind, &httpServer{
		logger:   logger,
		hostname: util.Hostname(),
		version:  util.Version(),
	})
	if err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(1)
	}
}
