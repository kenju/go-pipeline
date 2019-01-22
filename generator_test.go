package pipeline_test

import (
	"testing"
		"context"
	"github.com/kenju/go-pipeline"
	"reflect"
)

func TestGeneratorInterface(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []interface{}
	for v := range pipeline.GeneratorInterface(ctx, 1, 2, 3, 4) {
		results = append(results, v)
	}

	expected := []interface {}{1, 2, 3, 4}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestGeneratorByte(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []byte
	for v := range pipeline.GeneratorByte(ctx, 0x01, 0x02, 0x03) {
		results = append(results, v)
	}

	expected := []byte{0x01, 0x02, 0x03}
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

func TestGeneratorUint64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []uint64
	for v := range pipeline.GeneratorUint64(ctx, 1, 2, 3, 4) {
		results = append(results, v)
	}

	expected := []uint64{1, 2, 3, 4}
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

func TestGeneratorFloat64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float64
	for v := range pipeline.GeneratorFloat64(ctx, 1.1, 2.2, 3.3) {
		results = append(results, v)
	}

	expected := []float64{1.1, 2.2, 3.3}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
