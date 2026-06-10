module github.com/gost-dom/browser/internal/test/wpt

go 1.25.3

require (
	github.com/Marlliton/slogpretty v0.1.3
	github.com/gost-dom/browser v0.10.3
	github.com/gost-dom/browser/scripting/sobekengine v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.11.1
	golang.org/x/net v0.48.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4+incompatible // indirect
	github.com/google/pprof v0.0.0-20251114195745-4902fdda35c8 // indirect
	github.com/gost-dom/css v0.1.0 // indirect
	github.com/grafana/sobek v0.0.0-20251113105955-976a34df9c09 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gost-dom/browser => ../../../

replace github.com/gost-dom/browser/scripting/v8engine => ../../../scripting/v8engine/

replace github.com/gost-dom/browser/scripting/sobekengine => ../../../scripting/sobekengine/
