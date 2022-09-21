// generated-from:4397be19c8be568e4c5ad73f67dc7bead907d67838add9d87fc2799ede5e8c5c DO NOT REMOVE, DO UPDATE

package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/moov-io/base/log"
)

func NewTerminationListener() chan error {
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	return errs
}

func AwaitTermination(logger log.Logger, terminationListener chan error) error {
	if err := <-terminationListener; err != nil {
		return logger.Fatal().LogErrorf("Terminated: %v", err).Err()
	}
	return nil
}
