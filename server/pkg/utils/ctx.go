package utils

import "context"

// see https://eblog.fly.dev/backendbasics3.html -> #Building Client Middleware

type Key[T any] struct{} // key is a unique type that we can use as a key in a context

// WithValue returns a new context with the given value set.
// Only one value of each type can be set in a context;
// setting a value of the same type will overwrite the previous value.
func CtxWithValue[T any](ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, Key[T]{}, value)
}

// Value returns the value of type T in the given context,
// or false if the context does not contain a value of type T.
func CtxValue[T any](ctx context.Context) (T, bool) {
	value, ok := ctx.Value(Key[T]{}).(T)
	return value, ok
}
