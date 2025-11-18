package sobekhost

import (
	"fmt"
	"io"
	"strings"

	"github.com/gost-dom/browser/url"
	"github.com/grafana/sobek"
)

type sobekResolver struct {
	host    *gojaScriptHost
	ctx     *GojaContext
	modules map[sobek.ModuleRecord]string
	cache   map[string]sobek.ModuleRecord
}

func (m *sobekResolver) resolveModule(
	referencingScriptOrModule interface{},
	specifier string,
) (sobek.ModuleRecord, error) {
	ref := referencingScriptOrModule
	m.ctx.logger().
		Debug("SobekModule.ResolveModule", "ref", ref, "spec", specifier)
	var src string
	switch v := ref.(type) {
	case string:
		src = v
	case sobek.ModuleRecord:
		var ok bool
		src, ok = m.modules[v]
		if !ok {
			return nil, fmt.Errorf("ResolveModule: unknown source: %v", v)
		}
	default:
		return nil, fmt.Errorf("ResolveModule: ref not a string: (%T) %v", ref, ref)
	}
	name := url.ParseURLBase(specifier, src).Href()
	if cached, ok := m.cache[name]; ok {
		return cached, nil
	}
	code, err := m.download(name)
	if err != nil {
		return nil, err
	}
	m.ctx.logger().Debug("ParseModule", "name", name, "code", code)
	mod, err := sobek.ParseModule(name, code, m.resolveModule)
	if err == nil {
		m.ctx.logger().Info("Adding src", "mod", sobek.ModuleRecord(mod), "name", name)
		m.modules[mod] = name
		m.cache[name] = mod
	}
	if err == nil {
		err = mod.Link()
	}
	return mod, err
}

func (r *sobekResolver) download(url string) (string, error) {
	resp, err := r.host.HttpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("gost-dom/sobekhost: download errors: %w", err)
	}
	defer resp.Body.Close()
	var buf strings.Builder
	io.Copy(&buf, resp.Body)
	script := buf.String()

	if resp.StatusCode != 200 {
		err := fmt.Errorf(
			"gost-dom/sobekhost: ScriptContext: bad status code: %d, downloading %s",
			resp.StatusCode,
			url,
		)
		r.host.Logger.Error("Script download error", "err", err, "body", script)
		return "", err
	}
	return script, nil
}
