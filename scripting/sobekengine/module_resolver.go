package sobekengine

import (
	"fmt"

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
	referencingScriptOrModule interface{},
	specifier string,
) (sobek.ModuleRecord, error) {
	ref := referencingScriptOrModule
	m.ctx.logger().
		Info("SobekModule.ResolveModule", "ref", ref, "spec", specifier)
	var src *url.URL
	switch v := ref.(type) {
	case string:
		src = url.ParseURL(v)
	case sobek.ModuleRecord:
		var ok bool
		src, ok = m.modules[v]
		if !ok {
			return nil, fmt.Errorf("ResolveModule: unknown source: %v", v)
		}
	default:
		return nil, fmt.Errorf("ResolveModule: ref not a string: (%T) %v", ref, ref)
	}
	u := src.Join(specifier)
	href := u.Href()
	if cached, ok := m.cache[href]; ok {
		return cached, nil
	}
	code, err := gosthttp.Download(m.ctx.window.Context(), u, m.host.HttpClient)
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
