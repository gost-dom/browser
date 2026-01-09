// Package idlspec helps loading relevant specs from the webref module.
//
// This package may be solving two different problems.
//
//   - Gost-DOM only relies on a subset of webref specs.
//   - Exposed names are to be treated globally. When one package define an
//     interface or interface mixin with a specific name, this is the type
//     referenced in all mixins.
//
// An example for the latter. The fetch IDL specs defines "partial interface
// mixin WindowOrWorkerGlobalScope". As this is an "interface mixin", the
// operations should be installed on all classes that has this mixing. And as
// it's a partial, the information about the real name of the IDL interface that
// need to implement the interface is not present in the idl file.
//
// There is exactly one spec containing "interface mixin
// WindowOrWorkerGlobalScope", the HTML specs, which reveals that Window and
// WorkerGlobalScope uses the mixin.
//
// This package helps retrieve that information
package idlspec

import (
	"fmt"
	"iter"

	"github.com/gost-dom/code-gen/customrules"
	"github.com/gost-dom/webref/idl"
)

type spec struct {
	name    string
	idlSpec idl.Spec
	rules   customrules.SpecRules
}

type intf struct {
	spec string
	name string
	intf idl.Interface
}

var specs map[string]spec

var names map[string]intf

func IdlInterfaces() iter.Seq[idl.Interface] {
	return func(yield func(idl.Interface) bool) {
		for _, intf := range names {
			if !yield(intf.intf) {
				return
			}
		}
	}
}

func Interface(name string) (idl.Interface, bool) {
	res, ok := names[name]
	return res.intf, ok
}

func init() {
	specs = make(map[string]spec)
	names = make(map[string]intf)
	for _, name := range customrules.SpecNames() {
		rule := customrules.GetSpecRules(name)
		idlSpec, err := idl.Load(name)
		if err != nil {
			panic(fmt.Sprintf("Error loading spec: %v", err))
		}
		specs[name] = spec{name, idlSpec, rule}
		for intfName, i := range idlSpec.Interfaces {
			names[intfName] = intf{
				spec: name,
				name: intfName,
				intf: i,
			}
		}
	}
}
