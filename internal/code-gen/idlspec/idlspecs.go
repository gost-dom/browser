// Package idlspec helps loading relevant specs from the webref module.
//
// # This package may be solving two different
//
// - Gost-DOM doesn't rely on all webref specs.
// - Exposed names are to be treated globally. When one
//
// An example for the latter. The fetch IDL specs defines "partial interface
// mixin WindowOrWorkerGlobalScope". As this is an "interface mixin", the
// operations should be installed on all classes that has this mixing. And as
// it's a partial, that information is not present in this file.
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
	err     error
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

func init() {
	specs = make(map[string]spec)
	names = make(map[string]intf)
	for _, name := range customrules.SpecNames() {
		rule := customrules.GetSpecRules(name)
		idlSpec, err := idl.Load(name)
		if err != nil {
			panic(fmt.Sprintf("Error loading spec: %v", err))
		}
		specs[name] = spec{name, idlSpec, rule, err}
		for intfName, i := range idlSpec.Interfaces {
			names[intfName] = intf{
				spec: name,
				name: intfName,
				intf: i,
			}
		}
	}
}
