package main_test

import (
	"strings"
	"testing"

	. "github.com/gost-dom/browser/internal/test/wpt"

	"github.com/stretchr/testify/assert"
)

func TestManifestFile(t *testing.T) {
	// This isn't a _test_ - it was merely a tool to get feedback while
	// implementing JSON deserialization of the manifest
	data, err := LoadManifest()
	if !assert.NoError(t, err) {
		return
	}

	actual := make([]string, 100)
	i := 0
	for testCase := range data.All() {
		if !strings.HasPrefix(testCase.Path, "dom") {
			continue
		}
		actual[i] = testCase.Path
		i++
		if i >= 100 {
			break
		}
	}

	assert.Equal(t, expected, actual)
}

var expected = []string{
	"dom/abort/abort-signal-timeout.html",
	"dom/abort/reason-constructor.html",
	"dom/attributes-are-nodes.html",
	"dom/collections/HTMLCollection-as-prototype.html",
	"dom/collections/HTMLCollection-delete.html",
	"dom/collections/HTMLCollection-empty-name.html",
	"dom/collections/HTMLCollection-iterator.html",
	"dom/collections/HTMLCollection-own-props.html",
	"dom/collections/HTMLCollection-supported-property-indices.html",
	"dom/collections/HTMLCollection-supported-property-names.html",
	"dom/collections/domstringmap-supported-property-names.html",
	"dom/collections/namednodemap-supported-property-names.html",
	"dom/eventPathRemoved.html",
	"dom/events/Body-FrameSet-Event-Handlers.html",
	"dom/events/CustomEvent.html",
	"dom/events/Event-cancelBubble.html",
	"dom/events/Event-constants.html",
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
	"dom/events/EventListener-handleEvent-cross-realm.html",
	"dom/events/EventListener-handleEvent.html",
	"dom/events/EventListener-incumbent-global-1.sub.html",
	"dom/events/EventListener-incumbent-global-2.sub.html",
	"dom/events/EventListener-invoke-legacy.html",
	"dom/events/EventListenerOptions-capture.html",
	"dom/events/EventTarget-add-listener-platform-object.html",
	"dom/events/EventTarget-dispatchEvent-returnvalue.html",
	"dom/events/EventTarget-dispatchEvent.html",
	"dom/events/EventTarget-this-of-listener.html",
	"dom/events/KeyEvent-initKeyEvent.html",
	"dom/events/event-disabled-dynamic.html",
	"dom/events/event-global-is-still-set-when-coercing-beforeunload-result.html",
	"dom/events/event-global-is-still-set-when-reporting-exception-onerror.html",
	"dom/events/event-global.html",
	"dom/events/event-src-element-nullable.html",
	"dom/events/focus-event-document-move.html",
	"dom/events/handler-count.html",
	"dom/events/mouse-event-retarget.html",
	"dom/events/no-focus-events-at-clicking-editable-content-in-link.html",
	"dom/events/non-cancelable-when-passive/generic-events-stay-cancelable.html",
	"dom/events/non-cancelable-when-passive/non-passive-mousewheel-event-listener-on-body.html",
	"dom/events/non-cancelable-when-passive/non-passive-mousewheel-event-listener-on-div.html",
	"dom/events/non-cancelable-when-passive/non-passive-mousewheel-event-listener-on-document.html",
	"dom/events/non-cancelable-when-passive/non-passive-mousewheel-event-listener-on-root.html",
	"dom/events/non-cancelable-when-passive/non-passive-mousewheel-event-listener-on-window.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchmove-event-listener-on-body.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchmove-event-listener-on-div.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchmove-event-listener-on-document.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchmove-event-listener-on-root.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchmove-event-listener-on-window.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchstart-event-listener-on-body.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchstart-event-listener-on-div.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchstart-event-listener-on-document.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchstart-event-listener-on-root.html",
	"dom/events/non-cancelable-when-passive/non-passive-touchstart-event-listener-on-window.html",
	"dom/events/non-cancelable-when-passive/non-passive-wheel-event-listener-on-body.html",
	"dom/events/non-cancelable-when-passive/non-passive-wheel-event-listener-on-div.html",
	"dom/events/non-cancelable-when-passive/non-passive-wheel-event-listener-on-document.html",
	"dom/events/non-cancelable-when-passive/non-passive-wheel-event-listener-on-root.html",
	"dom/events/non-cancelable-when-passive/non-passive-wheel-event-listener-on-window.html",
	"dom/events/non-cancelable-when-passive/passive-mousewheel-event-listener-on-body.html",
	"dom/events/non-cancelable-when-passive/passive-mousewheel-event-listener-on-div.html",
	"dom/events/non-cancelable-when-passive/passive-mousewheel-event-listener-on-document.html",
	"dom/events/non-cancelable-when-passive/passive-mousewheel-event-listener-on-root.html",
}
