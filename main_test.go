package main

import (
	"fmt"
	"context"
	"testing"
	"math/rand"
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

func TestRepeatFn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	randFn := func() interface{} { return rand.Int() }
	for v := range Take(ctx, RepeatFn(ctx, randFn), 5) {
		fmt.Println(v)
	}
}
