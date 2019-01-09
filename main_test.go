package pipeline_test

import (
	"context"
	"fmt"
	"github.com/kenju/go-pipeline"
)

func Example() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for v := range pipeline.TakeInt(ctx, pipeline.GeneratorInt(ctx, 1, 2, 3, 4, 5), 3) {
		fmt.Printf("%v ", v)
	}
	// Output: 1 2 3
}

func Example_map() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mapFn := func(v string) string { return "**" + v + "**" }
	for v := range pipeline.MapString(ctx, mapFn, pipeline.GeneratorString(ctx, "hello", "world")) {
		fmt.Println(v)
	}
	// Output:
	// **hello**
	// **world**
}
