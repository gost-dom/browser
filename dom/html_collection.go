package dom

type HTMLCollection interface {
	All() []Element
	Length() int
	Item(int) Element
	NamedItem(string) Element
}

// htmlCollection provides an implementation for [HTMLCollection].
type htmlCollection struct{ node Node }

func newHtmlCollection(n Node) HTMLCollection { return htmlCollection{n} }

func (c htmlCollection) All() []Element {
	nodes := c.node.ChildNodes()
	l := nodes.Length()
	res := make([]Element, 0, l)
	for i := range l {
		n := nodes.Item(i)
		if e, ok := n.(Element); ok {
			res = append(res, e)
		}
	}
	return res
}

func (c htmlCollection) Length() int {
	return len(c.All())
}

func (c htmlCollection) Item(i int) Element {
	es := c.All()
	if i < 0 || i >= len(es) {
		return nil
	}
	return es[i]
}

func (c htmlCollection) NamedItem(name string) Element {
	for _, e := range c.All() {
		if id, hasId := e.GetAttribute("id"); hasId && id == name {
			return e
		}
		if n, hasName := e.GetAttribute("name"); hasName && n == name {
			return e
		}
	}
	return nil
}
