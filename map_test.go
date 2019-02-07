package pipeline_test

import (
	"context"
	"github.com/kenju/go-pipeline"
	"reflect"
	"testing"
)

func TestMapGenType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []pipeline.GenType
	mapFn := func(v pipeline.GenType) pipeline.GenType { return v }
	for v := range pipeline.MapGenType(ctx, mapFn, pipeline.GeneratorGenType(ctx, 1, 2, 3, 4)) {
		results = append(results, v)
	}

	expected := []pipeline.GenType{1, 2, 3, 4}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
