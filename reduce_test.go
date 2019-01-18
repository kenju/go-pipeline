package pipeline_test

import (
	"context"
	"github.com/kenju/go-pipeline"
	"math"
	"reflect"
	"testing"
)

func TestReduceInterface(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []interface{}
	reduceFn := func(v, acc interface{}) interface{} { return v }
	for v := range pipeline.ReduceInterface(ctx, reduceFn, pipeline.GeneratorInterface(ctx, 1, 2, 3, 4, 5)) {
		results = append(results, v)
	}

	expected := []interface{}{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestReduceInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	reduceFn := func(v, acc int) int { return v + acc }
	for v := range pipeline.ReduceInt(ctx, reduceFn, pipeline.GeneratorInt(ctx, 1, 2, 3, 4, 5)) {
		results = append(results, v)
	}

	expected := []int{1, 3, 6, 10, 15}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestReduceUint64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []uint64
	reduceFn := func(v, acc uint64) uint64 { return v + acc }
	for v := range pipeline.ReduceUint64(ctx, reduceFn, pipeline.GeneratorUint64(ctx, 1, 2, 3, 4, 5)) {
		results = append(results, v)
	}

	expected := []uint64{1, 3, 6, 10, 15}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestReduceString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []string
	reduceFn := func(v, acc string) string { return acc + "!!" + v }
	for v := range pipeline.ReduceString(ctx, reduceFn, pipeline.GeneratorString(ctx, "hello", "world")) {
		results = append(results, v)
	}

	expected := []string{"!!hello", "!!hello!!world"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestReduceFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float32
	reduceFn := func(v, acc float32) float32 { return v + acc }
	for v := range pipeline.ReduceFloat32(ctx, reduceFn, pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3)) {
		results = append(results, v)
	}

	float32AlmostEqual := func(a, b float32) bool {
		return math.Abs(float64(a)-float64(b)) <= 1e-6
	}

	expected := []float32{1.1, 3.3, 6.6}
	for i, v := range results {
		if !float32AlmostEqual(v, expected[i]) {
			t.Errorf("expected %.6f, got %.6f", expected[i], v)
		}
	}
}

func TestReduceFloat64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float64
	reduceFn := func(v, acc float64) float64 { return v + acc }
	for v := range pipeline.ReduceFloat64(ctx, reduceFn, pipeline.GeneratorFloat64(ctx, 1.1, 2.2, 3.3)) {
		results = append(results, v)
	}

	float64AlmostEqual := func(a, b float64) bool {
		return math.Abs(float64(a)-float64(b)) <= 1e-15
	}

	expected := []float64{1.1, 3.3, 6.6}
	for i, v := range results {
		if !float64AlmostEqual(v, expected[i]) {
			t.Errorf("expected %.6f, got %.6f", expected[i], v)
		}
	}
}
