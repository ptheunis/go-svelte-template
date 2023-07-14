package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/charmbracelet/log"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "⚡ ",
	})
)

func main() {
	httpsSrv, err := GetServer(8080)
	if err != nil {
		logger.Fatal(err)
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	go func() {

		if err := httpsSrv.ListenAndServe(); err != nil {
			logger.Fatalf("server failed with: %s", err)
		}
		/*
		   if err := httpsSrv.ListenAndServeTLS("", ""); err != nil {
		   			log.Fatalf("server failed with: %s", err)
		   		}
		*/
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	httpsSrv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Info("shutting down")

	os.Exit(0)
}
