package main

import (
	"fmt"
	"net/http"
    "os"
    "os/signal"
    "syscall"

	bookSvc "github.com/King-Rullmann-Book-Club/books-api/pkg/v1/service/books"
	"github.com/King-Rullmann-Book-Club/books-api/pkg/v1/transport/books"

    "github.com/go-kit/log"
)

func main() {
//	var (
//		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
//	)
//	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var s bookSvc.Service
	{
		s = bookSvc.NewService()
//		s = profilesvc.LoggingMiddleware(logger)(s)
	}

	var h http.Handler
	{
        h = books.NewTransport(s)
    }

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
        logger.Log("transport", "HTTP", "addr", ":8080")
        errs <- http.ListenAndServe(":8080", h)
	}()

	logger.Log("exit", <-errs)
}
