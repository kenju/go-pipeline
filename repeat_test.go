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
