package mutation_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/mutation"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	"github.com/stretchr/testify/assert"
)

func TestMutationObserverConformsToTheIDLSpecs(t *testing.T) {
	var observer any = &mutation.Observer{}

	_, observerConforms := observer.(dominterfaces.MutationObserver)
	assert.True(t, observerConforms,
		"mutation.Observer conforms to the MutationObserver IDL interface",
	)
}
