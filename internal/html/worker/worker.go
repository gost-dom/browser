package worker

import (
	"context"
	"errors"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
)

const queue_size = 32

// TODO: This should be WindowOrWorkerGlobalScope
type GlobalScope = htmlinterfaces.WorkerGlobalScope

type WorkItem func(GlobalScope) error

type queueItem struct {
	WorkItem
	done func()
}

type clockOption func(*Worker)

func WithContext(ctx context.Context) clockOption {
	return func(w *Worker) { w.ctx = ctx }
}

type Worker struct {
	queue  chan queueItem
	clock  *clock.Clock
	global *workerGlobalScope
	ctx    context.Context

	winClock       *clock.Clock
	winEventTarget event.EventTarget
}

// TODO: Implement global scope
func (w *Worker) scope() GlobalScope { return w.global }

func New(c *clock.Clock, opts ...clockOption) *Worker {
	res := &Worker{
		queue: make(chan queueItem, queue_size),
		clock: c,
	}
	res.global = newGlobal(res)
	for _, opt := range opts {
		opt(res)
	}
	if res.ctx != nil {
		context.AfterFunc(res.ctx, func() {
			res.Close()
		})
	}
	go func() {
		for item := range res.queue {
			item.WorkItem(res.scope())
			item.done()
		}
	}()
	return res
}

func FromWindow(win html.Window, c *clock.Clock) *Worker {
	w := New(c)
	w.winClock = win.Clock().(*clock.Clock)
	w.winEventTarget = win
	return w
}

func (w *Worker) Close() {
	if w.queue != nil {
		close(w.queue)
		w.queue = nil
		w.clock.Close()
	}
}

func (w Worker) Enqueue(item WorkItem) error {
	cb := w.clock.BeginEvent()
	cancel := func() { cb.AddEvent(func() error { return nil }) }
	enqueued := queueItem{
		item, cancel,
	}
	select {
	case w.queue <- enqueued:
		return nil
	default:
		cancel()
		return errors.New("gost-dom/worker: Queue full")
	}
}

func (w *Worker) Clock() *clock.Clock { return w.clock }
