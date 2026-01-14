package worker

import (
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

type Worker struct {
	queue  chan queueItem
	clock  *clock.Clock
	global *workerGlobalScope

	winClock       *clock.Clock
	winEventTarget event.EventTarget
}

// TODO: Implement global scope
func (w *Worker) scope() GlobalScope { return w.global }

func New(c *clock.Clock) *Worker {
	queue := make(chan queueItem, queue_size)
	res := &Worker{
		queue: queue,
		clock: c,
	}
	res.global = newGlobal(res)
	go func() {
		for item := range queue {
			item.WorkItem(res.scope())
			item.done()
		}
	}()
	return res
}

func FromWindow(win html.Window) *Worker {
	c := win.Clock().(*clock.Clock)
	w := New(c)
	w.winClock = c
	w.winEventTarget = win
	return w
}

func (w *Worker) Close() { close(w.queue) }

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
