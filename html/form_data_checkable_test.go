package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/assert"
)

// A checked checkbox contributes its value attribute. "on" is the default for
// a checkbox with no value, not a replacement for one -- without this, every
// checkbox in a group submits the same string and the group is unusable.
func TestFormDataCheckboxSubmitsItsValue(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input type="checkbox" name="mode" value="42" checked>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("42"), formData.Get("mode"))
}

// A valueless checkbox keeps the "on" default.
func TestFormDataCheckboxWithoutValueSubmitsOn(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input type="checkbox" name="agree" checked>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("on"), formData.Get("agree"))
}

// Several checked checkboxes sharing a name contribute one entry each --
// the shape a multi-select group relies on.
func TestFormDataCheckboxGroupSubmitsEveryCheckedValue(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input type="checkbox" name="mode" value="1" checked>
		<input type="checkbox" name="mode" value="2">
		<input type="checkbox" name="mode" value="3" checked>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t,
		[]html.FormDataValue{html.FormDataValue("1"), html.FormDataValue("3")},
		formData.GetAll("mode"),
		"only the checked boxes contribute, each with its own value")
}

// A radio group contributes only the checked option. Previously every radio
// was submitted regardless, so a group of three sent all three values.
func TestFormDataRadioSubmitsOnlyTheCheckedOption(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input type="radio" name="choice" value="a">
		<input type="radio" name="choice" value="b" checked>
		<input type="radio" name="choice" value="c">
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t,
		[]html.FormDataValue{html.FormDataValue("b")},
		formData.GetAll("choice"),
		"an unchecked radio contributes nothing")
}

// A radio group with nothing checked contributes nothing.
func TestFormDataRadioWithNoSelectionSubmitsNothing(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input name="username" value="john">
		<input type="radio" name="choice" value="a">
		<input type="radio" name="choice" value="b">
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.False(t, formData.Has("choice"))
	assert.Len(t, formData.Entries, 1)
}

// Checkedness falls back to the content attribute, so a control parsed from
// markup is seen as checked without anything calling SetChecked. This is what
// a test driving a rendered page relies on.
func TestFormDataCheckednessReadsTheAttribute(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input type="checkbox" name="cb" value="yes" checked>
		<input type="radio" name="r" value="picked" checked>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("yes"), formData.Get("cb"))
	assert.Equal(t, html.FormDataValue("picked"), formData.Get("r"))
}

// Disabled checkables are still excluded.
func TestFormDataSkipsDisabledCheckables(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input type="checkbox" name="cb" value="1" checked disabled>
		<input type="radio" name="r" value="a" checked disabled>
		<input name="username" value="john">
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.False(t, formData.Has("cb"))
	assert.False(t, formData.Has("r"))
	assert.Len(t, formData.Entries, 1)
}
