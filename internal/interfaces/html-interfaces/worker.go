package htmlinterfaces

type WorkerGlobalScope interface {
	WindowOrWorkerGlobalScope
	PostMessage(data any)
}
