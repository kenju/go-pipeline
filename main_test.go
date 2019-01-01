package pipeline

import (
	"fmt"
	"context"
	"testing"
	"math/rand"
	"time"
)

func TestGeneratorInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range GeneratorInt(ctx, 1, 2, 3, 4) {
		fmt.Printf("%v ", v)
	}
}

func TestAdd(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intCh := GeneratorInt(ctx, 1, 2, 3, 4)
	for v := range Add(ctx, intCh, 1) {
		fmt.Println(v)
	}

}

func TestMultiply(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intCh := GeneratorInt(ctx, 1, 2, 3, 4)
	for v := range Multiply(ctx, intCh, 10) {
		fmt.Println(v)
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

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	channels := []<-chan interface{}{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(1 * time.Second), // the close message of this channel will close the whole channels
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	<-Or(channels...)
	fmt.Printf("done after %v", time.Since(start))
}
