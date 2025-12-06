package fetch

import (
	"iter"
	"net/http"
	"slices"
	"strings"

	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/types"
)

// When iterating headers, have a max header count, to prevent overflow if
// client code continuously add new headers while iterating.
const MAX_HEADER_COUNT = 10000

type Headers struct {
	entity.Entity
	headers []Header
}

func parseHeaders(h http.Header) Headers {
	res := Headers{}
	for k, v := range h {
		for _, val := range v {
			// Cast the string values to ByteString without validation. This
			// assumes HTTP headers received from a server are valid headers.
			res.Append(types.ByteString(k), types.ByteString(val))
		}
	}
	return res
}

func compareHeaders(a, b Header) int {
	return strings.Compare(string(a.key), string(b.key))
}

func (h *Headers) Append(name, val types.ByteString) {
	// In order to have correct iteration behaviour, it's imperative that the
	// list is kept sorted.
	h.headers = insertSorted(h.headers, Header{key: name.ToLower(), val: val}, compareHeaders)
}

func (h *Headers) Delete(name types.ByteString) {
	name = name.ToLower()
	h.headers = slices.DeleteFunc(h.headers, func(h Header) bool { return h.key == name })
}

func (h *Headers) getRange(name types.ByteString) (first int, last int, found bool) {
	first = -1
	last = -1
	header := Header{key: name.ToLower()}
	first, found = slices.BinarySearchFunc(h.headers, header, compareHeaders)
	if !found {
		first = -1
		last = first
		return
	}
	l := len(h.headers)
	last = first
	for {
		last++
		if last >= l {
			return
		}
		if compareHeaders(header, h.headers[last]) < 0 {
			return
		}
	}
}

func (h *Headers) Get(name types.ByteString) (string, bool) {
	if h.headers == nil {
		return "", false
	}
	first, last, found := h.getRange(name)
	return h.formatValue(h.headers[first:last]), found
}

func (h *Headers) formatValue(val []Header) string {
	vv := make([]types.ByteString, len(val))
	for i, v := range val {
		vv[i] = v.val
	}
	return strings.Join(types.ByteStringsToStrings(vv), ", ")
}

func (h *Headers) Has(name types.ByteString) bool {
	_, ok := h.Get(name)
	return ok
}

func (h *Headers) Set(name, value types.ByteString) {
	name = name.ToLower()
	idx := slices.IndexFunc(h.headers, func(h Header) bool { return h.key == name })
	first, last, found := h.getRange(name)
	if found {
		h.headers = slices.Delete(h.headers, first+1, last)
		h.headers[idx].val = value
	} else {
		h.Append(name, value)
	}
}

// All() returns an iterator of key/value pairs of header keys and values.
// The same key may appear multiple times in the output. The returned order is
// guaranteed to be sorted by name. The iterator operates on a live list, i.e.
// new headers can be inserted while iterating. Panics if the number of headers
// iterated exceed MAX_HEADER_COUNT.
func (h *Headers) All() iter.Seq2[types.ByteString, types.ByteString] {
	return func(yield func(types.ByteString, types.ByteString) bool) {
		var collect []Header
		i := 0
		for i < len(h.headers) {
			assertHeaderCountWithinLimit(i)
			curr := h.headers[i]
			collect = append(collect, curr)
			if curr.key != "set-cookie" {
				var nextI = i + 1
				if len(h.headers) > nextI {
					if h.headers[nextI].key == curr.key {
						i++
						continue
					}
				}
			}
			if !yield(curr.key, types.ByteString(h.formatValue(collect))) {
				return
			}
			i++
			collect = nil
		}
	}
}
