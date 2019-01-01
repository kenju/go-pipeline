package pipeline_test

import (
	"testing"
	"fmt"
	"context"
	"github.com/kenju/go-pipeline"
)

func TestAdd(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intCh := pipeline.GeneratorInt(ctx, 1, 2, 3, 4)
	for v := range pipeline.Add(ctx, intCh, 1) {
		fmt.Println(v)
	}

}

func TestMultiply(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intCh := pipeline.GeneratorInt(ctx, 1, 2, 3, 4)
	for v := range pipeline.Multiply(ctx, intCh, 10) {
		fmt.Println(v)
	}

}
