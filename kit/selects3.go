package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	start := time.Now()
	tc := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop.")
			used := time.Since(start)
			fmt.Println(used)
			return
		default:
		}

		select {
		case <-tc.C:
			fmt.Println("in ticker.")
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	worker(ctx)
}
