package gosthttp

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gost-dom/browser/url"
)

// HttpDoer is the interface representing for the Do function on the http.Client
// interface.
type HttpDoer interface {
	Do(*http.Request) (*http.Response, error)
}

func Download(
	ctx context.Context,
	log *slog.Logger,
	url *url.URL,
	doer HttpDoer,
) (res string, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if log != nil {
		log.Info("download", slog.Any("url", url))
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		return
	}
	resp, err := doer.Do(req)
	if err != nil {
		return "", fmt.Errorf("gost-dom/gosthttp: download errors: %w", err)
	}
	defer resp.Body.Close()
	var buf strings.Builder
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return "", err
	}
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
