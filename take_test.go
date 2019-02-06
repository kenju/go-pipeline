package pipeline_test

import (
	"testing"
	"github.com/kenju/go-pipeline"
		"context"
	"reflect"
)

func TestTakeGenType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []pipeline.GenType
	for v := range pipeline.TakeGenType(ctx, pipeline.RepeatGenType(ctx, 1), 3) {
		results = append(results, v)
	}

	expected := []pipeline.GenType{1, 1, 1}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}

