package pipeline_test

import (
	"testing"
		"context"
	"github.com/kenju/go-pipeline"
	"reflect"
)

func TestRepeatGenType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []pipeline.GenType
	for v := range pipeline.TakeGenType(ctx, pipeline.RepeatGenType(ctx, 1), 10) {
		results = append(results, v)
	}

	expected := []pipeline.GenType{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

func TestRepeatFnGenType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []pipeline.GenType
	repeatFn := func() pipeline.GenType { return 3 }
	for v := range pipeline.TakeGenType(ctx, pipeline.RepeatFnGenType(ctx, repeatFn), 5) {
		results = append(results, v)
	}

	expected := []pipeline.GenType{3, 3, 3, 3, 3}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
