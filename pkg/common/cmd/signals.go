package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signals:
			cancel()
		case <-ctx.Done():
			return
		}

	}()

	return ctx
}
