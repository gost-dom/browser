package dom

// ElementOrDocument doesn't correspond to an interface in the web IDL specs. It
// contains common operations between dom Elements and documents.
type ElementOrDocument interface {
	GetElementsByTagName(string) NodeList
}

type elementOrDocument struct {
	node parentNode
}

func (e elementOrDocument) GetElementsByTagName(qualifiedName string) NodeList {
	res, err := e.node.QuerySelectorAll(qualifiedName)
	if err != nil {
		return &nodeList{}
	}
	return res
}
