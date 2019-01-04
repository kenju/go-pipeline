package pipeline_test

import (
	"testing"
	"fmt"
	"context"
	"github.com/kenju/go-pipeline"
)

func TestGeneratorInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.GeneratorInt(ctx, 1, 2, 3, 4) {
		fmt.Printf("%v ", v)
	}
}

func TestGeneratorString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.GeneratorString(ctx, "hello", "world", "!") {
		fmt.Printf("%v ", v)
	}
}

func TestGeneratorFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3) {
		fmt.Printf("%v ", v)
	}
}
