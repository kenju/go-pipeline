package main

import (
	"context"
)

func Repeat(
	ctx context.Context,
	values ...interface{},
) <-chan interface{} {
	valueCh := make(chan interface{})

	go func() {
		defer close(valueCh)

		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case valueCh <- v:
				}
			}
		}
	}()

	return valueCh
}

func RepeatFn(
	ctx context.Context,
	fn func() interface{},
) <-chan interface{} {
	valueCh := make(chan interface{})

	go func() {
		defer close(valueCh)

		for {
			select {
			case <-ctx.Done():
				return
			case valueCh <- fn():
			}
		}
	}()

	return valueCh
}

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
