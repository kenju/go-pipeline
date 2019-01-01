package main

import (
	"fmt"
	"context"
	"testing"
)

func TestRepeat(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for num := range Take(ctx, Repeat(ctx, 1), 10) {
		fmt.Printf("%v ", num)
	}
}
