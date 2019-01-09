package pipeline_test

import (
	"testing"
		"context"
	"github.com/kenju/go-pipeline"
	"reflect"
)

func TestRepeat(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []interface{}
	for v := range pipeline.TakeInterface(ctx, pipeline.RepeatInterface(ctx, 1), 10) {
		results = append(results, v)
	}

	expected := []interface{}{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []interface{}
	repeatFn := func() interface{} { return 3 }
	for v := range pipeline.TakeInterface(ctx, pipeline.RepeatFnInterface(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []interface{}{3, 3, 3, 3, 3}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []string
	for v := range pipeline.TakeString(ctx, pipeline.RepeatString(ctx, "hello"), 5) {
		results = append(results, v)
	}

	expected := []string{"hello", "hello", "hello", "hello", "hello"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFnString(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []string
	repeatFn := func() string { return "!" }
	for v := range pipeline.TakeString(ctx, pipeline.RepeatFnString(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []string{"!", "!", "!", "!", "!"}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	for v := range pipeline.TakeInt(ctx, pipeline.RepeatInt(ctx, 1), 10) {
		results = append(results, v)
	}

	expected := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFnInt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []int
	repeatFn := func() int { return 9 }
	for v := range pipeline.TakeInt(ctx, pipeline.RepeatFnInt(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []int{9, 9, 9, 9, 9}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatUint64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []uint64
	for v := range pipeline.TakeUint64(ctx, pipeline.RepeatUint64(ctx, 1), 10) {
		results = append(results, v)
	}

	expected := []uint64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFnUint64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []uint64
	repeatFn := func() uint64 { return 9 }
	for v := range pipeline.TakeUint64(ctx, pipeline.RepeatFnUint64(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []uint64{9, 9, 9, 9, 9}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float32
	for v := range pipeline.TakeFloat32(ctx, pipeline.RepeatFloat32(ctx, 1.1), 5) {
		results = append(results, v)
	}

	expected := []float32{1.1, 1.1, 1.1, 1.1, 1.1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFnFloat32(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float32
	repeatFn := func() float32 { return 1.1 }
	for v := range pipeline.TakeFloat32(ctx, pipeline.RepeatFnFloat32(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []float32{1.1, 1.1, 1.1, 1.1, 1.1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFloat64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float64
	for v := range pipeline.TakeFloat64(ctx, pipeline.RepeatFloat64(ctx, 1.1), 5) {
		results = append(results, v)
	}

	expected := []float64{1.1, 1.1, 1.1, 1.1, 1.1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFnFloat64(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []float64
	repeatFn := func() float64 { return 1.1 }
	for v := range pipeline.TakeFloat64(ctx, pipeline.RepeatFnFloat64(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []float64{1.1, 1.1, 1.1, 1.1, 1.1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
