package main_test

import (
	"bytes"
	_ "embed"
	"io"
	"log/slog"
	"strings"
	"testing"

	. "github.com/gost-dom/browser/internal/test/wpt"

	"github.com/stretchr/testify/assert"
)

//go:embed manifest.json
var manifest []byte

func TestManifestFile2(t *testing.T) {
	actual := make([]string, 100)
	i := 0
	nullLog := slog.New(slog.NewTextHandler(io.Discard, nil))
	for testCase := range ParseManifest(t.Context(), bytes.NewReader(manifest), nullLog) {
		if !strings.HasPrefix(testCase.Path, "dom") {
			continue
		}
		actual[i] = testCase.Path
		i++
		if i >= 100 {
			break
		}
	}

	assert.Equal(t, expected2, actual)
}

var expected2 = []string{
	"dom/abort/AbortSignal.any.html",
	"dom/abort/AbortSignal.any.shadowrealm-in-dedicatedworker.html",
	"dom/abort/AbortSignal.any.shadowrealm-in-shadowrealm.html",
	"dom/abort/AbortSignal.any.shadowrealm-in-sharedworker.html",
	"dom/abort/AbortSignal.any.shadowrealm-in-window.html",
	"dom/abort/AbortSignal.any.worker.html",
	"dom/abort/AbortSignal.https.any.shadowrealm-in-audioworklet.html",
	"dom/abort/AbortSignal.https.any.shadowrealm-in-serviceworker.html",
	"dom/abort/abort-signal-any.any.html",
	"dom/abort/abort-signal-any.any.worker.html",
	"dom/abort/abort-signal-timeout.html",
	"dom/abort/event.any.html",
	"dom/abort/event.any.shadowrealm-in-dedicatedworker.html",
	"dom/abort/event.any.shadowrealm-in-shadowrealm.html",
	"dom/abort/event.any.shadowrealm-in-sharedworker.html",
	"dom/abort/event.any.shadowrealm-in-window.html",
	"dom/abort/event.any.worker.html",
	"dom/abort/event.https.any.shadowrealm-in-audioworklet.html",
	"dom/abort/event.https.any.shadowrealm-in-serviceworker.html",
	"dom/abort/reason-constructor.html",
	"dom/abort/timeout-shadowrealm.any.shadowrealm-in-dedicatedworker.html",
	"dom/abort/timeout-shadowrealm.any.shadowrealm-in-shadowrealm.html",
	"dom/abort/timeout-shadowrealm.any.shadowrealm-in-sharedworker.html",
	"dom/abort/timeout-shadowrealm.any.shadowrealm-in-window.html",
	"dom/abort/timeout-shadowrealm.https.any.shadowrealm-in-audioworklet.html",
	"dom/abort/timeout-shadowrealm.https.any.shadowrealm-in-serviceworker.html",
	"dom/abort/timeout.any.html",
	"dom/abort/timeout.any.worker.html",
	"dom/attributes-are-nodes.html",
	"dom/collections/HTMLCollection-as-prototype.html",
	"dom/collections/HTMLCollection-delete.html",
	"dom/collections/HTMLCollection-empty-name.html",
	"dom/collections/HTMLCollection-iterator.html",
	"dom/collections/HTMLCollection-live-mutations.window.html",
	"dom/collections/HTMLCollection-own-props.html",
	"dom/collections/HTMLCollection-supported-property-indices.html",
	"dom/collections/HTMLCollection-supported-property-names.html",
	"dom/collections/domstringmap-supported-property-names.html",
	"dom/collections/namednodemap-supported-property-names.html",
	"dom/eventPathRemoved.html",
	"dom/events/AddEventListenerOptions-once.any.html",
	"dom/events/AddEventListenerOptions-once.any.worker.html",
	"dom/events/AddEventListenerOptions-passive.any.html",
	"dom/events/AddEventListenerOptions-passive.any.worker.html",
	"dom/events/AddEventListenerOptions-signal.any.html",
	"dom/events/AddEventListenerOptions-signal.any.worker.html",
	"dom/events/Body-FrameSet-Event-Handlers.html",
	"dom/events/CustomEvent.html",
	"dom/events/Event-cancelBubble.html",
	"dom/events/Event-constants.html",
	"dom/events/Event-constructors.any.html",
	"dom/events/Event-constructors.any.worker.html",
	"dom/events/Event-defaultPrevented-after-dispatch.html",
	"dom/events/Event-defaultPrevented.html",
	"dom/events/Event-dispatch-bubble-canceled.html",
	"dom/events/Event-dispatch-bubbles-false.html",
	"dom/events/Event-dispatch-bubbles-true.html",
	"dom/events/Event-dispatch-click.html",
	"dom/events/Event-dispatch-click.tentative.html",
	"dom/events/Event-dispatch-detached-click.html",
	"dom/events/Event-dispatch-detached-input-and-change.html",
	"dom/events/Event-dispatch-handlers-changed.html",
	"dom/events/Event-dispatch-listener-order.window.html",
	"dom/events/Event-dispatch-multiple-cancelBubble.html",
	"dom/events/Event-dispatch-multiple-stopPropagation.html",
	"dom/events/Event-dispatch-omitted-capture.html",
	"dom/events/Event-dispatch-on-disabled-elements.html",
	"dom/events/Event-dispatch-order-at-target.html",
	"dom/events/Event-dispatch-order.html",
	"dom/events/Event-dispatch-other-document.html",
	"dom/events/Event-dispatch-propagation-stopped.html",
	"dom/events/Event-dispatch-redispatch.html",
	"dom/events/Event-dispatch-reenter.html",
	"dom/events/Event-dispatch-single-activation-behavior.html",
	"dom/events/Event-dispatch-target-moved.html",
	"dom/events/Event-dispatch-target-removed.html",
	"dom/events/Event-dispatch-throwing-multiple-globals.html",
	"dom/events/Event-dispatch-throwing.html",
	"dom/events/Event-init-while-dispatching.html",
	"dom/events/Event-initEvent.html",
	"dom/events/Event-isTrusted.any.html",
	"dom/events/Event-isTrusted.any.worker.html",
	"dom/events/Event-propagation.html",
	"dom/events/Event-returnValue.html",
	"dom/events/Event-stopImmediatePropagation.html",
	"dom/events/Event-stopPropagation-cancel-bubbling.html",
	"dom/events/Event-subclasses-constructors.html",
	"dom/events/Event-timestamp-cross-realm-getter.html",
	"dom/events/Event-timestamp-high-resolution.html",
	"dom/events/Event-timestamp-high-resolution.https.html",
	"dom/events/Event-timestamp-safe-resolution.html",
	"dom/events/Event-type-empty.html",
	"dom/events/Event-type.html",
	"dom/events/EventListener-addEventListener.sub.window.html",
	"dom/events/EventListener-handleEvent-cross-realm.html",
	"dom/events/EventListener-handleEvent.html",
	"dom/events/EventListener-incumbent-global-1.sub.html",
	"dom/events/EventListener-incumbent-global-2.sub.html",
	"dom/events/EventListener-invoke-legacy.html",
	"dom/events/EventListenerOptions-capture.html",
}
