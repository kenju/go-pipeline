package pipeline_test

import (
	"testing"
	"github.com/kenju/go-pipeline"
		"context"
	"reflect"
)

func TestTake(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []interface{}
	for v := range pipeline.Take(ctx, pipeline.Repeat(ctx, 1), 3) {
		results = append(results, v)
	}

	expected := []interface{}{1, 1, 1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestTakeInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	for v := range pipeline.TakeInt(ctx, pipeline.GeneratorInt(ctx, 1, 2, 3), 1) {
		results = append(results, v)
	}

	expected := []int{1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestTakeString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []string
	for v := range pipeline.TakeString(ctx, pipeline.GeneratorString(ctx, "hello", "world", "!"), 1) {
		results = append(results, v)
	}

	expected := []string{"hello"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestTakeFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float32
	for v := range pipeline.TakeFloat32(ctx, pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3), 1) {
		results = append(results, v)
	}

	expected := []float32{1.1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
