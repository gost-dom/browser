package worker

import (
	"time"

	"github.com/gost-dom/browser/internal/clock"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
)

var _ htmlinterfaces.WindowOrWorkerGlobalScope = &workerGlobalScope{}

type workerGlobalScope struct{ worker *Worker }

func newGlobal(w *Worker) *workerGlobalScope { return &workerGlobalScope{worker: w} }

func (g *workerGlobalScope) SetTimeout(
	cb clock.TaskCallback,
	delay time.Duration,
) clock.TaskHandle {
	handle := g.worker.clock.SetTimeout(
		func() error {
			g.worker.Enqueue(func(GlobalScope) error {
				return cb()
			})
			return nil
		}, delay)
	return handle
}

func (g *workerGlobalScope) ClearTimeout(clock.TaskHandle) {
	panic("Not implemented")
}

func (g *workerGlobalScope) SetInterval(clock.TaskCallback, time.Duration) clock.TaskHandle {
	panic("Not implemented")
}

func (g *workerGlobalScope) ClearInterval(clock.TaskHandle) {
	panic("Not implemented")
}

func (g *workerGlobalScope) QueueMicrotask(clock.TaskCallback) {
	panic("Not implemented")
}
