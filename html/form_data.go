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
	inputs := form.Elements()
	formData := NewFormData()
	for _, input := range inputs.All() {
		if input, ok := input.(HTMLInputElement); ok {
			name := input.Name()
			if name == "" {
				continue
			}
			switch input.Type() {
			case "submit":
				continue
			case "checkbox":
				if input.Checked() {
					formData.Append(name, "on")
				}
			default:
				// TODO: handle no values
				formData.Append(name, NewFormDataValueString(input.Value()))
			}
		}
	}
	return formData
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
