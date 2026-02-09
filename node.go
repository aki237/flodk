package flodk

import "context"

// Node represents any node of the execution graph.
type Node[T any] interface {
	// Execute runs the node logic with the passed in state.
	Execute(ctx context.Context, state T) (T, error)
}

// FunctionNode is a function type which implements the Node interface.
// This is useful when the [Node]s don't require any preloaded state.
type FunctionNode[T any] func(ctx context.Context, state T) (T, error)

// Execute implements the [Node] interface for FunctionNode
func (fn FunctionNode[T]) Execute(ctx context.Context, state T) (T, error) {
	return fn(ctx, state)
}

// noop is a Node which does nothing
type noop[T any] struct{}

// Execute implements the [Node] interface for [noop]
func (n noop[T]) Execute(ctx context.Context, state T) (T, error) {
	return state, nil
}

// Noop returns a [Node] which does nothing.
func Noop[T any]() Node[T] {
	return noop[T]{}
}

// ConditionalNode is a node type which inspects the state to
// decide the next [Node].
type ConditionalNode[T any] interface {
	Execute(ctx context.Context, state T) string
}

// ConditionalFunction is a function type which implements the ConditionalNode interface.
// This is used when no pre-defined state is necessary for the conditional node.
type ConditionalFunction[T any] func(ctx context.Context, state T) string

// Execute implements the [ConditionalNode] interface for FunctionNode
func (fn ConditionalFunction[T]) Execute(ctx context.Context, state T) string {
	return fn(ctx, state)
}
