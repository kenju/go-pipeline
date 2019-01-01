package main

import (
	"fmt"
	"context"
	"testing"
)

func TestGeneratorInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range GeneratorInt(ctx, 1, 2, 3, 4) {
		fmt.Printf("%v ", v)
	}
}

func TestRepeat(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range Take(ctx, Repeat(ctx, 1), 10) {
		fmt.Printf("%v ", v)
	}
}
