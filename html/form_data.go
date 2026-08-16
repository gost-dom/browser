package html

import (
	"io"
	"iter"
	"slices"
	"strings"

	netURL "net/url"

	"github.com/gost-dom/browser/dom"
)

type FormDataValue string // TODO Blob/file

func NewFormDataValueString(value string) FormDataValue { return FormDataValue(value) }

type FormDataEntry struct {
	Name  string
	Value FormDataValue
}

type FormData struct {
	Entries []FormDataEntry
}

func NewFormData() *FormData {
	return &FormData{nil}
}

// NewFormDataForm creates a new FormData initialised with values from a form.
//
// see also: https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#constructing-the-form-data-set
func NewFormDataForm(form HTMLFormElement) *FormData {
	elements := form.Elements()
	formData := NewFormData()
	for el := range elements.All() {
		if domEl, ok := el.(dom.Element); ok && isDisabledControl(domEl) {
			continue
		}
		if input, ok := el.(HTMLInputElement); ok {
			name := input.Name()
			if name == "" {
				continue
			}
			switch input.Type() {
			case "submit":
				continue
			case "checkbox", "radio":
				// Only checked controls contribute, and they contribute their
				// value attribute -- "on" is merely the default when no value
				// was given, not a replacement for one.
				//
				// see https://html.spec.whatwg.org/multipage/input.html#checkbox-state-(type=checkbox):concept-fe-value
				if input.Checked() {
					value := input.Value()
					if value == "" {
						value = "on"
					}
					formData.Append(name, NewFormDataValueString(value))
				}
			default:
				// TODO: handle no values
				formData.Append(name, NewFormDataValueString(input.Value()))
			}
			continue
		}
		// <select> and <textarea> have no specialized DOM types yet, so read
		// their name/value from the element directly.
		if domEl, ok := el.(dom.Element); ok {
			name, hasName := domEl.GetAttribute("name")
			if !hasName || name == "" {
				continue
			}
			switch domEl.TagName() {
			case "SELECT":
				for _, v := range selectValues(domEl) {
					formData.Append(name, NewFormDataValueString(v))
				}
			case "TEXTAREA":
				formData.Append(name, NewFormDataValueString(domEl.TextContent()))
			}
		}
	}
	return formData
}

// isDisabledControl reports whether a form control is barred from the form
// data set because it is disabled -- either by its own attribute, or by
// sitting inside a disabled <fieldset>.
//
// The fieldset rule has an exception: controls inside that fieldset's *first*
// <legend> stay enabled, so a disabled fieldset can still carry a working
// control in its caption.
//
// see https://html.spec.whatwg.org/multipage/form-control-infrastructure.html#constructing-the-form-data-set
// and https://html.spec.whatwg.org/multipage/form-elements.html#concept-fieldset-disabled
func isDisabledControl(el dom.Element) bool {
	if _, disabled := el.GetAttribute("disabled"); disabled {
		return true
	}
	fieldset, err := el.Closest("fieldset[disabled]")
	if err != nil || fieldset == nil {
		return false
	}
	// Inside a disabled fieldset. Exempt only if the control sits within that
	// fieldset's first legend child.
	legend := firstLegendChild(fieldset)
	if legend == nil {
		return true
	}
	for anc := el.ParentElement(); anc != nil; anc = anc.ParentElement() {
		if anc == legend {
			return false
		}
		if anc == fieldset {
			break
		}
	}
	return true
}

// firstLegendChild returns a fieldset's first direct <legend> child, or nil.
func firstLegendChild(fieldset dom.Element) dom.Element {
	parent, ok := fieldset.(dom.ParentNode)
	if !ok {
		return nil
	}
	for _, child := range parent.Children().All() {
		if child.TagName() == "LEGEND" {
			return child
		}
	}
	return nil
}

// selectValues returns the values a <select> contributes to the form data set:
// one entry per selected <option>, in document order.
//
//   - <select multiple> contributes every selected option, and nothing at all
//     when none are selected.
//   - A single select contributes exactly one value; with no option selected it
//     falls back to the first option, which is what a browser submits for a
//     select the user never touched.
//   - A select with no options contributes nothing either way.
//
// Disabled options, and options inside a disabled <optgroup>, are skipped
// entirely -- they can be neither selected nor used as the fallback.
//
// An option with no "value" attribute contributes its text content with ASCII
// whitespace stripped and collapsed, per
// https://html.spec.whatwg.org/multipage/form-elements.html#concept-option-value
func selectValues(sel dom.Element) []string {
	parent, ok := sel.(dom.ParentNode)
	if !ok {
		return nil
	}
	options, err := parent.QuerySelectorAll("option")
	if err != nil {
		return nil
	}

	_, multiple := sel.GetAttribute("multiple")

	var selected []string
	var firstValue string
	var hasFirst bool
	for i := 0; i < options.Length(); i++ {
		opt, ok := options.Item(i).(dom.Element)
		if !ok {
			continue
		}
		if isDisabledOption(opt) {
			continue
		}
		val, hasVal := opt.GetAttribute("value")
		if !hasVal {
			val = collapseASCIIWhitespace(opt.TextContent())
		}
		if !hasFirst {
			firstValue, hasFirst = val, true
		}
		if _, isSelected := opt.GetAttribute("selected"); isSelected {
			if !multiple {
				return []string{val}
			}
			selected = append(selected, val)
		}
	}
	if multiple {
		// No fallback: an untouched multiple select submits nothing.
		return selected
	}
	if !hasFirst {
		return nil // no options at all, or all of them disabled
	}
	return []string{firstValue}
}

// isDisabledOption reports whether an <option> is unselectable: disabled
// itself, or inside a disabled <optgroup>.
func isDisabledOption(opt dom.Element) bool {
	if _, disabled := opt.GetAttribute("disabled"); disabled {
		return true
	}
	group, err := opt.Closest("optgroup[disabled]")
	return err == nil && group != nil
}

// collapseASCIIWhitespace strips leading/trailing ASCII whitespace and
// collapses interior runs to a single space, which is what an <option>'s
// text-derived value is normalised to. Templates indent their options, so
// without this a value arrives wrapped in the markup's newlines and tabs.
func collapseASCIIWhitespace(s string) string {
	return strings.Join(strings.FieldsFunc(s, isASCIIWhitespace), " ")
}

func isASCIIWhitespace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r', '\f':
		return true
	}
	return false
}

func (d *FormData) AddElement(e dom.Element) {
	if name, _ := e.GetAttribute("name"); name != "" {
		value, _ := e.GetAttribute("value")
		d.Append(name, NewFormDataValueString(value))
	}
}

func (d *FormData) Append(name string, value FormDataValue) {
	d.Entries = append(d.Entries, FormDataEntry{name, value})
}

func (d *FormData) Set(name string, value FormDataValue) {
	predicate := elementByName(name)
	i := slices.IndexFunc(d.Entries, predicate)
	if i == -1 {
		d.Append(name, value)
		return
	} else {
		d.Delete(name)
		d.Entries = slices.Insert(d.Entries, i, FormDataEntry{
			Name:  name,
			Value: value,
		})
	}
}

func (d *FormData) Keys() []string {
	result := make([]string, len(d.Entries))
	for i, e := range d.Entries {
		result[i] = e.Name
	}
	return result
}

func (d *FormData) Values() []FormDataValue {
	result := make([]FormDataValue, len(d.Entries))
	for i, e := range d.Entries {
		result[i] = e.Value
	}
	return result
}

func (d *FormData) Delete(name string) {
	d.Entries = slices.DeleteFunc(
		d.Entries,
		elementByName(name),
	)
}

func (d *FormData) Get(name string) FormDataValue {
	for _, e := range d.Entries {
		if e.Name == name {
			return e.Value
		}
	}
	return ""
}

func (d *FormData) GetAll(name string) []FormDataValue {
	var result []FormDataValue
	for _, e := range d.Entries {
		if e.Name == name {
			result = append(result, e.Value)
		}
	}
	return result
}

func (d *FormData) Has(name string) bool {
	return slices.IndexFunc(d.Entries, elementByName(name)) != -1
}

func (d *FormData) GetReader() io.ReadCloser {
	return io.NopCloser(strings.NewReader(d.QueryString()))
}

// QueryString returns the formdata as a &-separated URL encoded key-value pair.
func (d *FormData) QueryString() string {
	sb := strings.Builder{}
	for i, e := range d.Entries {
		if i != 0 {
			sb.WriteString("&")
		}

		sb.WriteString(netURL.QueryEscape(e.Name))
		sb.WriteString("=")
		sb.WriteString(netURL.QueryEscape(string(e.Value)))
	}
	return sb.String()
}

func (d *FormData) All() iter.Seq2[string, FormDataValue] {
	return func(yield func(string, FormDataValue) bool) {
		for _, e := range d.Entries {
			if !yield(e.Name, e.Value) {
				return
			}
		}
	}
}

type Predicate[T any] func(T) bool

func elementByName(name string) Predicate[FormDataEntry] {
	return func(e FormDataEntry) bool { return e.Name == name }
}
