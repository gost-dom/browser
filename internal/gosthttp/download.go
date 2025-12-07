package gosthttp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HttpDoer is the interface representing for the Do function on the http.Client
// interface.
type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

func Download(ctx context.Context, url string, doer HttpDoer) (res string, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	resp, err := doer.Do(req)
	if err != nil {
		return "", fmt.Errorf("gost-dom/gosthttp: download errors: %w", err)
	}
	defer resp.Body.Close()
	var buf strings.Builder
	io.Copy(&buf, resp.Body)
	script := buf.String()

	if resp.StatusCode != 200 {
		err = fmt.Errorf(
			"gost-dom/gosthttp: bad status code: %d, downloading %s",
			resp.StatusCode,
			url,
		)
	}
	return script, err

}
