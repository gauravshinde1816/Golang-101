package main

import (
	"context"
	"fmt"
	"time"
)

func addParamsToContext(ctx context.Context) context.Context {

	return context.WithValue(ctx, "api-key", "cguyqwgcqwhgui")
}

func doSomethingWithContext(ctx context.Context) {

	for {

		select {
		case <-ctx.Done():
			fmt.Println("timed out")
		default:
			fmt.Println("doing something cool")
			time.Sleep(500 * time.Millisecond)

		}

	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = addParamsToContext(ctx)

	go doSomethingWithContext(ctx)

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Deadline has passsed")
		fmt.Println(err)
	}

}
