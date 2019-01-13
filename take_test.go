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
	for v := range pipeline.TakeInterface(ctx, pipeline.RepeatInterface(ctx, 1), 3) {
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

func TestTakeUint64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []uint64
	for v := range pipeline.TakeUint64(ctx, pipeline.GeneratorUint64(ctx, 1, 2, 3), 1) {
		results = append(results, v)
	}

	expected := []uint64{1}
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

func TestTakeFloat64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float64
	for v := range pipeline.TakeFloat64(ctx, pipeline.GeneratorFloat64(ctx, 1.1, 2.2, 3.3), 1) {
		results = append(results, v)
	}

	expected := []float64{1.1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}