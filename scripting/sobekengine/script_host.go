package sobekengine

import (
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/cache"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/grafana/sobek"
)

const internal_symbol_name = "__go_dom_internal_value__"

// Deprecated: Don't create script hosts directly, create them through an
// engine.
func New() html.ScriptHost {
	return &scriptHost{}
}

type scriptHost struct {
	httpClient  *http.Client
	logger      *slog.Logger
	initializer *internal.ScriptEngineConfigurer[jsTypeParam]
	cache       *cache.Cache
	clock       *clock.Clock
}

type propertyNameMapper struct{}

func (_ propertyNameMapper) FieldName(t reflect.Type, f reflect.StructField) string {
	return ""
}

func uncapitalize(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

func (_ propertyNameMapper) MethodName(t reflect.Type, m reflect.Method) string {
	var doc dom.Document
	var document = reflect.TypeOf(&doc).Elem()
	if t.Implements(document) && m.Name == "Location" {
		return uncapitalize(m.Name)
	} else {
		return ""
	}
}

func (h *scriptHost) NewContext(bc html.BrowsingContext) html.ScriptContext {
	vm := sobek.New()
	vm.SetFieldNameMapper(propertyNameMapper{})
	result := &scriptContext{
		host:         h,
		cache:        h.cache,
		vm:           vm,
		browsingCtx:  bc,
		wrappedGoObj: sobek.NewSymbol(internal_symbol_name),
		classes:      make(map[string]*class),
	}

	globalThis := vm.GlobalObject()
	globalThis.DefineDataPropertySymbol(
		result.wrappedGoObj,
		vm.ToValue(bc),
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
	)
	h.initializer.Configure(result)
	if global := result.global; global != nil {
		globalThis.SetPrototype(result.global.prototype)
	}
	if e, ok := bc.(entity.Components); ok {
		codec.SetJsValue(e, newScope(result).GlobalThis())
	}

	return result
}

func (d *scriptHost) Close() {}
