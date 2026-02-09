package flodk

import "context"

// EdgeResolver interface defines the edge routing configuration
// which returns the next node id for the passed state.
type EdgeResolver[T any] interface {
	Resolve(ctx context.Context, state T) string
}

// ConstEdge is simple implementation of the EdgeResolver which
// returns a constant next node id no matter what the current the state.
type ConstEdge[T any] string

// Resolve implements the [EdgeResolver] which is used to resolve to a constant
// node regardless of the current state.
func (c ConstEdge[T]) Resolve(ctx context.Context, state T) string {
	return string(c)
}

// ConditionalEdge is used to redirect to different branches based on the
// value returned by the [ConditionalNode].
type ConditionalEdge[T any] struct {
	exec         ConditionalNode[T]
	redirections map[string]string
}

// Resolve implements the [EdgeResolver] interface for [ConditionalEdge].
func (ce ConditionalEdge[T]) Resolve(ctx context.Context, state T) string {
	next, ok := ce.redirections[ce.exec.Execute(ctx, state)]
	if !ok {
		return ""
	}

	return next
}
