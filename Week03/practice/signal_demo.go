package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

const shortDuration = 50 * time.Second

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, os.Kill)

	select {
	case <-time.After(30 * time.Second):
		fmt.Println("overslept")
	case s := <-signalCh:
		fmt.Println("Got signal:", s)
	case <-ctx.Done():
		fmt.Printf("ctx err: %v", ctx.Err())
	}
}
