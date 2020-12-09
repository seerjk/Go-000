package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 5 * time.Second

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Printf("ctx err: %v", ctx.Err())
	}
}
