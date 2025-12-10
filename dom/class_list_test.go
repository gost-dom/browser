package dom_test

import (
	"testing"

	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/stretchr/testify/assert"

	"github.com/onsi/gomega"
)

func initClassListPair(doc Document) (res struct {
	element   Element
	classList DOMTokenList
}) {
	res.element = doc.CreateElement("div")
	res.classList = res.element.ClassList()
	return
}

func initElement() Element {
	doc := CreateHTMLDocument()
	return doc.CreateElement("div")
}

func TestClassList(t *testing.T) {
	t.Parallel()

	expect := gomega.NewWithT(t).Expect
	doc := CreateHTMLDocument()
	{
		pair := initClassListPair(doc)
		expect(pair.classList.Add("c1", "c2")).To(Succeed())
		expect(pair.element).To(
			HaveAttribute("class", "c1 c2"), "Multiple classes added are separated by space")
	}
	{
		pair := initClassListPair(doc)
		pair.element.SetAttribute("class", "c1 c2")
		expect(pair.classList.Add("c1", "c3")).To(Succeed())
		expect(pair.element).To(
			HaveAttribute("class", "c1 c2 c3"), "Adding the same class twice is ignored")
	}
	{
		pair := initClassListPair(doc)
		pair.element.SetAttribute("class", "c1 c2")
		expect(pair.classList.Add("c1", "c3")).To(Succeed())
		expect(pair.element).To(
			HaveAttribute("class", "c1 c2 c3"), "Adding the same class twice is ignored")
	}
	{
		pair := initClassListPair(doc)
		err := pair.classList.Add("")
		assert.ErrorIs(t, err, ErrSyntax, "Adding empty string to class list is SyntaxError")
		assert.NotErrorIs(t, err, ErrInvalidCharacter)

		err = pair.classList.Add("a b")
		assert.NotErrorIs(t, err, ErrSyntax)
		assert.ErrorIs(t, err, ErrInvalidCharacter,
			"Adding empty string to class list is an InvalidCharacterException",
		)
		assert.ErrorIs(t, err, ErrDom, "SyntaxError is a DOMException")
	}
}

func TestClassListAll(t *testing.T) {
	doc := CreateHTMLDocument()
	pair := initClassListPair(doc)
	pair.element.SetAttribute("class", "a b c")
	expect(t, pair.classList.All()).To(ConsistOf("a", "b", "c"))
}

func TestClassListContains(t *testing.T) {
	doc := CreateHTMLDocument()
	pair := initClassListPair(doc)
	pair.element.SetAttribute("class", "a b c")
	assert.True(t, pair.classList.Contains("a"), "Contains an existing class")
	assert.False(t, pair.classList.Contains("x"), "Contains a non-existing class")
}

func TestClassListLength(t *testing.T) {
	doc := CreateHTMLDocument()
	pair := initClassListPair(doc)
	pair.element.SetAttribute("class", "a b c")
	assert.Equal(t, 3, pair.classList.Length())
}

func TestClassListValueReflectClassAttribute(t *testing.T) {
	doc := CreateHTMLDocument()
	el := doc.CreateElement("div")
	el.SetAttribute("class", "a b c")
	assert.Equal(t, "a b c", el.ClassList().Value(), "Reading class list")

	el.ClassList().SetValue("a b  z")
	got, _ := el.GetAttribute("class")
	assert.Equal(t, "a b  z", got)
}

func TestClassListItem(t *testing.T) {
	doc := CreateHTMLDocument()
	el := doc.CreateElement("div")
	el.SetAttribute("class", "a b c")
	item1, foundItem1 := el.ClassList().Item(1)
	assert.True(t, foundItem1)
	expect(t, item1).To(Equal("b"), "Reading indexed element")
	_, foundItem3 := el.ClassList().Item(3)
	assert.False(t, foundItem3, "Reading indexed out of range should return nil")
}

func TestClassListRemove(t *testing.T) {
	el := initElement()
	el.SetAttribute("class", "a b c")
	el.ClassList().Remove("b")
	expect(t, el).To(HaveAttribute("class", "a c"), "Remove existing class")

	el.SetAttribute("class", "a b c")
	el.ClassList().Remove("c")
	expect(t, el).To(HaveAttribute("class", "a b"), "Remove the last class")

	el.SetAttribute("class", "a b c")
	el.ClassList().Remove("x")
	expect(t, el).To(HaveAttribute("class", "a b c"), "Remove non-existing class")
}

func TestClassListReplace(t *testing.T) {
	el := initElement()
	el.SetAttribute("class", "a b c")
	expect(
		t,
		el.ClassList().Replace("b", "x"),
	).To(BeTrue(), "Replace return value for existing class")
	expect(t, el).To(HaveAttribute("class", "a c x"), "Class attribute it updated")

	expect(
		t,
		el.ClassList().Replace("y", "x"),
	).To(BeFalse(), "Replace return value for non-existing class")
	expect(t, el).To(HaveAttribute("class", "a c x"), "Class attribute is preserved")
}
