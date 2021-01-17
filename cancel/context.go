package cancel

import (
	"context"
	"time"
)

type Context interface {
	context.Context
	Cancel()
}

type cancelCtx struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// Creates a context that wraps the given context and returns an obj that can be cancelled.
// This allows an object which desires to cancel a long running operation to store a single
// cancel.Context in it's struct variables instead of having to store both the context.Context
// and context.CancelFunc.
func New(ctx context.Context) Context {
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithCancel(ctx)
	return &cancelCtx{
		cancel: cancel,
		ctx:    ctx,
	}
}

func (c *cancelCtx) Cancel()                                 { c.cancel() }
func (c *cancelCtx) Deadline() (deadline time.Time, ok bool) { return c.ctx.Deadline() }
func (c *cancelCtx) Done() <-chan struct{}                   { return c.ctx.Done() }
func (c *cancelCtx) Err() error                              { return c.ctx.Err() }
func (c *cancelCtx) Value(key interface{}) interface{}       { return c.ctx.Value(key) }
