package pipeline_test

import (
	"testing"
	"time"
	"fmt"
	"github.com/kenju/go-pipeline"
)

func ExampleOr() {
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

	<-pipeline.OrInterface(channels...)
	fmt.Printf("done after %v", time.Since(start))
}

func ExampleOrString() {
	sig := func(after time.Duration) <-chan string {
		c := make(chan string)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	channels := []<-chan string{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(1 * time.Second), // the close message of this channel will close the whole channels
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	<-pipeline.OrString(channels...)
	fmt.Printf("done after %v", time.Since(start))
}

func ExampleOrInt() {
	sig := func(after time.Duration) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	channels := []<-chan int{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(1 * time.Second), // the close message of this channel will close the whole channels
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	<-pipeline.OrInt(channels...)
	fmt.Printf("done after %v", time.Since(start))
}

func ExampleOrFloat32() {
	sig := func(after time.Duration) <-chan float32 {
		c := make(chan float32)
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	channels := []<-chan float32{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(1 * time.Second), // the close message of this channel will close the whole channels
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	<-pipeline.OrFloat32(channels...)
	fmt.Printf("done after %v", time.Since(start))
}
