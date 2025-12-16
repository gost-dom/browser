package dom

type predicate[T any] interface {
	Match(T) bool
}

type predicateFunc[T any] func(T) bool

func (f predicateFunc[T]) Match(t T) bool { return f(t) }

// liveHtmlCollection implements interface HTML collection and delivers a live
// collection of a subtree mathing a specific filter.
type liveHtmlCollection struct {
	nodeRev int
	root    Element
	cache   []*liveHtmlCollection
	items   []Element
	filter  predicate[Element]
}

func newLiveHtmlCollection(n Node, filter predicate[Element]) *liveHtmlCollection {
	switch t := n.(type) {
	case Element:
		return &liveHtmlCollection{root: t, filter: filter}
	case Document:
		return &liveHtmlCollection{root: t.DocumentElement(), filter: filter}
	}
	panic("Bad node")
}

var _ HTMLCollection = liveHtmlCollection{}

func (c *liveHtmlCollection) checkCache() {
	nodeRev := c.root.revision()
	if nodeRev == c.nodeRev {
		return
	}
	c.nodeRev = nodeRev

	children := c.root.Children().All()
	c.items = make([]Element, 0)
	c.cache = make([]*liveHtmlCollection, len(children))
	if c.filter.Match(c.root) {
		c.items = append(c.items, c.root)
	}
	for i, child := range children {
		subItem := &liveHtmlCollection{root: child, filter: c.filter}
		subItem.checkCache()
		c.cache[i] = subItem
		c.items = append(c.items, subItem.items...)
	}
}

func (c liveHtmlCollection) All() []Element { c.checkCache(); return c.items }
func (c liveHtmlCollection) Length() int    { c.checkCache(); return len(c.items) }
func (c liveHtmlCollection) Item(i int) Element {
	c.checkCache()
	if i < 0 || i >= len(c.items) {
		return nil
	}
	return c.items[i]
}
func (c liveHtmlCollection) NamedItem(string) Element { panic("NamedItem: Not implemented") }
