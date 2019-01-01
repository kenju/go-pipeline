package pipeline_test

import (
	"testing"
	"github.com/kenju/go-pipeline"
	"fmt"
	"context"
)

func TestTake(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.Take(ctx, pipeline.Repeat(ctx, 1), 3) {
		fmt.Printf("%v ", v)
	}
}
