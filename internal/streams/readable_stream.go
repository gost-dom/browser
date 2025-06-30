package streams

import "github.com/gost-dom/browser/internal/promise"

type GetReaderOption func()

type ReadableStream interface {
	GetReader(...GetReaderOption) Reader
}

// ReadResult represents the result of calling the browser function [ReadableStream.read]
//
// [ReadableStream.read]: https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamDefaultReader/read
type ReadResult struct {
	Value []byte
	Done  bool
}

type Reader interface {
	Read() promise.Promise[ReadResult]
}

type ReadableStreamDefaultReader interface{ Reader }
