package pipeline_test

import (
	"testing"
	"context"
	"github.com/kenju/go-pipeline"
	"reflect"
	"math"
)

func TestMapInterface(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []interface{}
	mapFn := func(v interface{}) interface {} { return v }
	for v := range pipeline.MapInterface(ctx, mapFn, pipeline.GeneratorInterface(ctx, 1, 2, 3, 4)) {
		results = append(results, v)
	}

	expected := []interface {}{1, 2, 3, 4}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestMapInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	mapFn := func(v int) int { return v + 1 }
	for v := range pipeline.MapInt(ctx, mapFn, pipeline.GeneratorInt(ctx, 1, 2, 3, 4)) {
		results = append(results, v)
	}

	expected := []int{2, 3, 4, 5}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestMapString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []string
	mapFn := func(v string) string { return "**" + v + "**" }
	for v := range pipeline.MapString(ctx, mapFn, pipeline.GeneratorString(ctx, "hello", "world")) {
		results = append(results, v)
	}

	expected := []string{"**hello**", "**world**"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestMapFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float32
	mapFn := func(v float32) float32 { return v + 0.1 }
	for v := range pipeline.MapFloat32(ctx, mapFn, pipeline.GeneratorFloat32(ctx, 1.1, 2.2, 3.3)) {
		results = append(results, v)
	}

	float32AlmostEqual := func(a, b float32) bool {
		return math.Abs(float64(a)-float64(b)) <= 1e-1
	}

	expected := []float32{1.2, 2.3, 3.3}
	for i, v := range results {
		if !float32AlmostEqual(v, expected[i]) {
			t.Errorf("expected %.6f, got %.6f", expected[i], v)
		}
	}
}
