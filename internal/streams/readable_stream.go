package streams

import "github.com/gost-dom/browser/internal/promise"

type GetReaderOption func()

type ReadableStream interface {
	GetReader(...GetReaderOption) Reader
}

type ReadResult struct{}

type Reader interface {
	Read() promise.Promise[ReadResult]
}

type ReadableStreamDefaultReader interface{ Reader }
