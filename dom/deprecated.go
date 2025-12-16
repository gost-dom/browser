package dom

// Deprecated: This will be removed
type Attributes []Attr

func (attrs Attributes) Length() int {
	return len(attrs)
}
