package dom

// htmlCollection provides an implementation for [HTMLCollection].
type htmlCollection struct{ node Node }

func newHtmlCollection(n Node) HTMLCollection { return htmlCollection{n} }

func (c htmlCollection) elements() []Element {
	nodes := c.node.ChildNodes().All()
	res := make([]Element, 0, len(nodes))
	for _, n := range nodes {
		if e, ok := n.(Element); ok {
			res = append(res, e)
		}
	}
	return res
}

func (c htmlCollection) Length() int {
	return len(c.elements())
}

func (c htmlCollection) Item(i int) Element {
	es := c.elements()
	if i < 0 || i >= len(es) {
		return nil
	}
	return es[i]
}

func (c htmlCollection) NamedItem(name string) Element {
	for _, e := range c.elements() {
		if id, hasId := e.GetAttribute("id"); hasId && id == name {
			return e
		}
		if n, hasName := e.GetAttribute("name"); hasName && n == name {
			return e
		}
	}
	return nil
}
