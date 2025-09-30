package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
}

func parse(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):

		case <-ctx.Done():
			fmt.Println("deadline exceeded")
			return
		}
	}
}
