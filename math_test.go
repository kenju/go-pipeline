package pipeline_test

import (
	"testing"
		"context"
	"github.com/kenju/go-pipeline"
	"reflect"
)

func TestAdd(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	intCh := pipeline.GeneratorInt(ctx, 1, 2, 3, 4)
	for v := range pipeline.Add(ctx, intCh, 1) {
		results = append(results, v)
	}

	expected := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestMultiply(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	intCh := pipeline.GeneratorInt(ctx, 1, 2, 3, 4)
	for v := range pipeline.Multiply(ctx, intCh, 10) {
		results = append(results, v)
	}

	expected := []int{10, 20, 30, 40}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
