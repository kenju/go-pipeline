package pipeline_test

import (
	"testing"
		"context"
	"github.com/kenju/go-pipeline"
	"reflect"
)

func TestGeneratorGenType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var results []pipeline.GenType
	for v := range pipeline.GeneratorGenType(ctx, 1, 2, 3, 4) {
		results = append(results, v)
	}

	expected := []pipeline.GenType {1, 2, 3, 4}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
