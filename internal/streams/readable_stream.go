package streams

type GetReaderOption func()

type ReadableStream interface {
	GetReader(...GetReaderOption) Reader
}

type Reader interface{}
