package pipeline_test

import (
	"context"
	"github.com/kenju/go-pipeline"
	"reflect"
	"testing"
)

func TestReduceGenType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []pipeline.GenType
	reduceFn := func(v, acc pipeline.GenType) pipeline.GenType { return v }
	for v := range pipeline.ReduceGenType(ctx, reduceFn, pipeline.GeneratorGenType(ctx, 1, 2, 3, 4, 5)) {
		results = append(results, v)
	}

	expected := []pipeline.GenType{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
