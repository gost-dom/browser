package fetch

import "io"

type Body interface {
	io.Reader
}
