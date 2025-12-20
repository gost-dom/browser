package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/suite"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

func TestFormData(t *testing.T) {
	t.Parallel()

	Expect(t, NewFormData()).To(BeEmptyFormData(), "New form data is empty")

	t.Run("Append multiple values with same key", func(t *testing.T) {
		suite.Run(t, new(FormDataMultipleValuesTestSuite))
	})
}

type FormDataMultipleValuesTestSuite struct {
	gosttest.GomegaSuite
	formData *FormData
}

func (s *FormDataMultipleValuesTestSuite) SetupTest() {
	s.formData = NewFormData()
	s.formData.Append("Key1", "Value1")
	s.formData.Append("Key2", "Value2")
	s.formData.Append("Key1", "Value3")
	s.formData.Append("Key3", "Value4")
}

func (s *FormDataMultipleValuesTestSuite) TestQueryFunctions() {
	s.Expect(
		s.formData).To(
		HaveEntries(
			"Key1", "Value1",
			"Key2", "Value2",
			"Key1", "Value3",
			"Key3", "Value4",
		),
		"FormData.Entries() returns all entries, including duplicate keys",
	)

	s.Expect(
		s.formData.Keys()).To(
		HaveExactElements(
			"Key1",
			"Key2",
			"Key1",
			"Key3",
		),
		"FormData.Keys() return all keys, including duplicates",
	)
	s.Expect(
		s.formData.Values()).To(
		HaveExactElements(
			BeFormDataStringValue("Value1"),
			BeFormDataStringValue("Value2"),
			BeFormDataStringValue("Value3"),
			BeFormDataStringValue("Value4"),
		),
		"FormData.Values() return all values, including duplicate keys",
	)
	s.Expect(
		s.formData.Get("Key1")).To(
		BeEquivalentTo("Value1"),
		"Get returns the first value for duplicate keys",
	)

	s.Expect(
		s.formData.GetAll("Key1")).To(
		HaveExactElements(
			BeFormDataStringValue("Value1"),
			BeFormDataStringValue("Value3"),
		),
		"Get all returns all values for duplicate keys",
	)

	s.Expect(
		s.formData.Has("Key1")).To(
		BeTrue(),
		"Has() called with existing key",
	)

	s.Expect(
		s.formData.Has("non-existing-key")).To(
		BeFalse(),
		"Has() called with non-existing key",
	)
}

func (s *FormDataMultipleValuesTestSuite) TestRemoveValue() {
	s.formData.Delete("Key1")
	s.Expect(s.formData).To(HaveEntries(
		"Key2", "Value2",
		"Key3", "Value4"))
}

func (s *FormDataMultipleValuesTestSuite) TestSetValueWithExistingName() {
	s.formData.Set("Key1", "Value5")
	s.Expect(s.formData).To(HaveEntries(
		"Key1", "Value5",
		"Key2", "Value2",
		"Key3", "Value4"))
}

func (s *FormDataMultipleValuesTestSuite) TestSetValueWithNewName() {
	s.formData.Set("Key4", "Value5")
	s.Expect(s.formData).To(HaveEntries(
		"Key1", "Value1",
		"Key2", "Value2",
		"Key1", "Value3",
		"Key3", "Value4",
		"Key4", "Value5"))
}

func TestFormDataFormWithUnnamedInput(t *testing.T) {
	doc := htmltest.NewHTMLDocumentHelper(t, nil)
	form := NewHtmlFormElement(doc.HTMLDocument)
	Expect := gomega.NewWithT(t).Expect

	namedInput1 := NewHTMLInputElement(doc)
	namedInput1.SetName("username")
	namedInput1.SetValue("john")

	unnamedInput := NewHTMLInputElement(doc)
	unnamedInput.SetValue("should-be-skipped")

	namedInput2 := NewHTMLInputElement(doc)
	namedInput2.SetName("email")
	namedInput2.SetValue("john@example.com")

	form.Append(namedInput1, unnamedInput, namedInput2)
	formData := html.NewFormDataForm(form)

	Expect(formData.Has("username")).To(BeTrue(), "Form data includes named input 'username'")
	Expect(formData.Get("username")).
		To(Equal(html.FormDataValue("john")), "Form data has correct value for 'username'")
	Expect(formData.Has("email")).To(BeTrue(), "Form data includes named input 'email'")
	Expect(formData.Get("email")).
		To(Equal(html.FormDataValue("john@example.com")), "Form data has correct value for 'email'")
	Expect(len(formData.Entries)).
		To(Equal(2), "Form data has exactly 2 entries, skipping the unnamed input")
}

func BeEmptyFormData() types.GomegaMatcher {
	return gcustom.MakeMatcher(func(data *FormData) (bool, error) {
		return len(data.Entries) == 0, nil
	})
}

func HaveEntries(entries ...string) types.GomegaMatcher {
	if len(entries)%2 != 0 {
		panic("Wrong number of entries, must be even")
	}
	noOfEntries := len(entries) / 2
	expected := make([]FormDataEntry, noOfEntries)
	for i := range noOfEntries {
		j := i * 2
		expected[i] = FormDataEntry{
			Name:  entries[j],
			Value: FormDataValue(entries[j+1]),
		}
	}
	return WithTransform(
		func(data *FormData) []FormDataEntry { return data.Entries },
		HaveExactElements(expected),
	)
}

func BeFormDataStringValue(value string) types.GomegaMatcher {
	return BeEquivalentTo(value)
}
