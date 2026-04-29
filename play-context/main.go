package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func longRunningProcess(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("Failed: Process took too long")
}

func ghibliReq(ctx context.Context, requestString string) {
	req, _ := http.NewRequest(http.MethodGet, requestString, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("request failed", err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Response received, status code: ", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("took too long")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := longRunningProcess(ctx)
		if err != nil {
			cancel()
		}
	}()

	ghibliReq(ctx, "https://ghibliapi.herokuapp.com/films/8611129-2dbc-4a81-a72f-77ddfc1b1b49")
}
