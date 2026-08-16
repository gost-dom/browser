package html_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// formFromHTML loads a document and returns its first <form>.
func formFromHTML(t *testing.T, src string) html.HTMLFormElement {
	t.Helper()
	win := htmltest.NewWindowHelper(t, nil)
	win.MustLoadHTML(src)
	el, err := win.Document().QuerySelector("form")
	require.NoError(t, err, "querying for form")
	require.NotNil(t, el, "no form in test HTML")
	form, ok := el.(html.HTMLFormElement)
	require.True(t, ok, "element is not an HTMLFormElement")
	return form
}

// The form data set must include <select> and <textarea> controls, not just
// <input>. https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#constructing-the-form-data-set
func TestFormDataFormIncludesSelectAndTextarea(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input name="username" value="john">
		<select name="colour">
			<option value="red">Red</option>
			<option value="green" selected>Green</option>
			<option value="blue">Blue</option>
		</select>
		<textarea name="comment">hello world</textarea>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("john"), formData.Get("username"))
	assert.Equal(t, html.FormDataValue("green"), formData.Get("colour"),
		"select submits the option carrying the selected attribute")
	assert.Equal(t, html.FormDataValue("hello world"), formData.Get("comment"),
		"textarea submits its text content")
}

// With no option marked selected, a browser submits the first one.
func TestFormDataFormSelectDefaultsToFirstOption(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour">
			<option value="red">Red</option>
			<option value="green">Green</option>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("red"), formData.Get("colour"))
}

// An option without a value attribute submits its text content.
func TestFormDataFormSelectOptionWithoutValueUsesText(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour">
			<option>Red</option>
			<option selected>Green</option>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("Green"), formData.Get("colour"))
}

// A <select multiple> contributes one entry per selected option, in document
// order -- not just the first.
func TestFormDataFormMultipleSelectSubmitsEverySelectedOption(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour" multiple>
			<option value="red" selected>Red</option>
			<option value="green">Green</option>
			<option value="blue" selected>Blue</option>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t,
		[]html.FormDataValue{html.FormDataValue("red"), html.FormDataValue("blue")},
		formData.GetAll("colour"),
		"every selected option is submitted, in document order")
}

// A <select multiple> with nothing selected submits nothing -- the
// first-option fallback applies only to single selects.
func TestFormDataFormMultipleSelectWithNoSelectionSubmitsNothing(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input name="username" value="john">
		<select name="colour" multiple>
			<option value="red">Red</option>
			<option value="green">Green</option>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.False(t, formData.Has("colour"),
		"an untouched multiple select contributes no entry")
	assert.Len(t, formData.Entries, 1, "only the input should be submitted")
}

// Malformed markup: a non-multiple <select> cannot legally have two selected
// options, but nothing stops an author writing it. Browsers honour the first
// and ignore the rest; we do the same rather than emitting two entries for a
// control that can only hold one value.
func TestFormDataFormSingleSelectWithTwoSelectedOptionsTakesTheFirst(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour">
			<option value="red" selected>Red</option>
			<option value="green" selected>Green</option>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t,
		[]html.FormDataValue{html.FormDataValue("red")},
		formData.GetAll("colour"),
		"a non-multiple select contributes exactly one value, the first selected")
}

// A select with no options contributes nothing, rather than an empty value.
func TestFormDataFormSelectWithNoOptionsSubmitsNothing(t *testing.T) {
	for _, tc := range []struct {
		name string
		src  string
	}{
		{"single", `<form><input name="username" value="john"><select name="colour"></select></form>`},
		{"multiple", `<form><input name="username" value="john"><select name="colour" multiple></select></form>`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			formData := html.NewFormDataForm(formFromHTML(t, tc.src))
			assert.False(t, formData.Has("colour"),
				"a select with no options contributes no entry")
			assert.Len(t, formData.Entries, 1)
		})
	}
}

// An option's text-derived value is stripped and collapsed of ASCII
// whitespace -- templates indent their options, and the raw text content
// would otherwise arrive wrapped in newlines and tabs.
func TestFormDataFormSelectCollapsesOptionTextWhitespace(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour">
			<option>
				Forest   Green
			</option>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("Forest Green"), formData.Get("colour"),
		"leading/trailing whitespace stripped, interior runs collapsed to one space")
}

// A disabled control is barred from the form data set.
func TestFormDataFormSkipsDisabledControls(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input name="username" value="john">
		<input name="nickname" value="johnny" disabled>
		<select name="colour" disabled><option value="red" selected>Red</option></select>
		<textarea name="comment" disabled>hello</textarea>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.True(t, formData.Has("username"))
	assert.False(t, formData.Has("nickname"), "disabled input is skipped")
	assert.False(t, formData.Has("colour"), "disabled select is skipped")
	assert.False(t, formData.Has("comment"), "disabled textarea is skipped")
	assert.Len(t, formData.Entries, 1)
}

// Controls inside a disabled <fieldset> are disabled too -- except those in
// the fieldset's first <legend>.
func TestFormDataFormSkipsControlsInDisabledFieldset(t *testing.T) {
	form := formFromHTML(t, `<form>
		<fieldset disabled>
			<legend><input name="in_legend" value="kept"></legend>
			<input name="in_fieldset" value="dropped">
			<select name="sel_in_fieldset"><option value="x" selected>X</option></select>
		</fieldset>
		<input name="outside" value="kept">
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.True(t, formData.Has("outside"), "control outside the fieldset is unaffected")
	assert.True(t, formData.Has("in_legend"), "control in the first legend stays enabled")
	assert.False(t, formData.Has("in_fieldset"), "control in a disabled fieldset is skipped")
	assert.False(t, formData.Has("sel_in_fieldset"), "select in a disabled fieldset is skipped")
}

// An enabled fieldset disables nothing.
func TestFormDataFormEnabledFieldsetKeepsControls(t *testing.T) {
	form := formFromHTML(t, `<form>
		<fieldset>
			<legend>Details</legend>
			<input name="username" value="john">
		</fieldset>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("john"), formData.Get("username"))
}

// Disabled options can be neither selected nor used as the fallback.
func TestFormDataFormSkipsDisabledOptions(t *testing.T) {
	t.Run("disabled option is not the fallback", func(t *testing.T) {
		form := formFromHTML(t, `<form>
			<select name="colour">
				<option value="placeholder" disabled>Choose...</option>
				<option value="red">Red</option>
			</select>
		</form>`)

		formData := html.NewFormDataForm(form)

		assert.Equal(t, html.FormDataValue("red"), formData.Get("colour"),
			"the disabled placeholder is skipped, so red is the first usable option")
	})

	t.Run("disabled option is not selected", func(t *testing.T) {
		form := formFromHTML(t, `<form>
			<select name="colour" multiple>
				<option value="red" selected>Red</option>
				<option value="green" selected disabled>Green</option>
			</select>
		</form>`)

		formData := html.NewFormDataForm(form)

		assert.Equal(t,
			[]html.FormDataValue{html.FormDataValue("red")},
			formData.GetAll("colour"),
			"a disabled option contributes nothing even when marked selected")
	})

	t.Run("all options disabled contributes nothing", func(t *testing.T) {
		form := formFromHTML(t, `<form>
			<input name="username" value="john">
			<select name="colour"><option value="red" disabled>Red</option></select>
		</form>`)

		formData := html.NewFormDataForm(form)

		assert.False(t, formData.Has("colour"))
		assert.Len(t, formData.Entries, 1)
	})
}

// Options inside a disabled <optgroup> are disabled too.
func TestFormDataFormSkipsOptionsInDisabledOptgroup(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour">
			<optgroup label="Discontinued" disabled>
				<option value="puce">Puce</option>
			</optgroup>
			<optgroup label="Current">
				<option value="red">Red</option>
			</optgroup>
		</select>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("red"), formData.Get("colour"),
		"options in a disabled optgroup are skipped, including as the fallback")
}

// Unnamed controls are skipped, matching the existing behaviour for inputs.
func TestFormDataFormSkipsUnnamedSelectAndTextarea(t *testing.T) {
	form := formFromHTML(t, `<form>
		<input name="username" value="john">
		<select><option value="red" selected>Red</option></select>
		<textarea>orphan text</textarea>
	</form>`)

	formData := html.NewFormDataForm(form)

	assert.Len(t, formData.Entries, 1, "only the named input should be submitted")
	assert.Equal(t, html.FormDataValue("john"), formData.Get("username"))
}

// Setting the selected attribute after parse -- what a test driving a page
// does -- must change what the form submits.
func TestFormDataFormSelectReflectsAttributeChange(t *testing.T) {
	form := formFromHTML(t, `<form>
		<select name="colour">
			<option value="red">Red</option>
			<option value="green">Green</option>
		</select>
	</form>`)

	opts, err := form.QuerySelectorAll("option")
	require.NoError(t, err)
	green, ok := opts.Item(1).(dom.Element)
	require.True(t, ok)
	green.SetAttribute("selected", "selected")

	formData := html.NewFormDataForm(form)

	assert.Equal(t, html.FormDataValue("green"), formData.Get("colour"))
}
