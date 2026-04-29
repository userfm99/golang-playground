package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals

		cancel()
	}()

	wg.Add(1)
	go func() {
		if err := myWorker(ctx); err != nil {
			cancel()
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if err := startServer(ctx); err != nil {
			cancel()
		}
		wg.Done()
	}()

	wg.Wait()
}

func myWorker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := doSomethingRepeatedly()
			if err != nil {
				return err
			}
		}
	}
}

func startServer(ctx context.Context) error {

	var srv http.Server

	srv.Addr = ":8080"

	go func() {
		<-ctx.Done() // Wait for the context to be done

		// Shutdown the server
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		return fmt.Errorf("HTTP server ListenAndServe: %w", err)
	}

	return nil
}

func doSomethingRepeatedly() error {
	return nil
}
