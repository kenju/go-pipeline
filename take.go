package main

import (
	"context"
)

func Take(
	ctx context.Context,
	valueCh <-chan interface{},
	num int,
) <-chan interface{} {
	takeCh := make(chan interface{})

	go func() {
		defer close(takeCh)

		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case takeCh <- <-valueCh:
			}
		}
	}()

	return takeCh
}