package pipeline_test

import (
	"testing"
		"context"
	"github.com/kenju/go-pipeline"
	"reflect"
)

func TestGeneratorInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	for v := range pipeline.GeneratorInt(ctx, 1, 2, 3, 4) {
		results = append(results, v)
	}

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestGeneratorString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []string
	for v := range pipeline.GeneratorString(ctx, "hello", "world", "!") {
		results = append(results, v)
	}

	expected := []string{"hello", "world", "!"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestGeneratorFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float32
	for v := range pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3) {
		results = append(results, v)
	}

	expected := []float32{1.1, 2.2, 3.3}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
