package sobekengine

import (
	"fmt"
	"log/slog"

	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/url"
	"github.com/grafana/sobek"
)

type moduleResolver struct {
	host    *scriptHost
	ctx     *scriptContext
	modules map[sobek.ModuleRecord]*url.URL
	cache   map[string]sobek.ModuleRecord
}

func (m *moduleResolver) resolveModule(
	ref interface{},
	specifier string,
) (sobek.ModuleRecord, error) {
	m.ctx.logger().Info("SobekModule.ResolveModule",
		slog.Any("ref", ref),
		slog.String("spec", specifier))

	if v, ok := ref.(sobek.ModuleRecord); ok {
		if src, ok := m.modules[v]; ok {
			return m.resolveModuleUrl(src.Join(specifier))
		}
		return nil, fmt.Errorf("ResolveModule: unknown source: %v", v)
	} else {
		return nil, fmt.Errorf("ResolveModule: unexpected ref type: (%T) %v", ref, ref)
	}
}

func (m *moduleResolver) resolveModuleUrl(
	u *url.URL,
) (sobek.ModuleRecord, error) {
	href := u.Href()
	if cached, ok := m.cache[href]; ok {
		return cached, nil
	}
	code, err := gosthttp.Download(m.ctx.Context(), u, m.host.httpClient)
	if err != nil {
		return nil, err
	}
	mod, err := sobek.ParseModule(href, code, m.resolveModule)
	if err == nil {
		m.modules[mod] = u
		m.cache[href] = mod
	}
	return mod, err
}
