package gosthttp

import (
	"context"
	"net/http"
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestHttpHandlerCreateContext(t *testing.T) {
	t.Parallel()
	var ctx context.Context
	server := func(w http.ResponseWriter, r *http.Request) {
		ctx = r.Context()
	}

	client := NewHttpClientFromHandler(http.HandlerFunc(server))
	r, err := client.Get("http://localhost/foo")

	g := gomega.NewWithT(t)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(r.StatusCode).To(Equal(200))

	g.Expect(ctx).To(Equal(context.Background()))
}
