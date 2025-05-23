package dom

import (
	"iter"
	"slices"
	"strings"
)

type DOMTokenList struct {
	attribute string
	element   Element
}

func NewClassList(element Element) DOMTokenList {
	return DOMTokenList{"class", element}
}

func NewDOMTokenList(attribute string, element Element) DOMTokenList {
	return DOMTokenList{attribute, element}
}

func (l DOMTokenList) Add(tokens ...string) error {
	tokenList := l.getTokens()
	for _, token := range tokens {
		if token == "" {
			return newSyntaxError("token is empty")
		}
		if strings.Contains(token, " ") {
			return newInvalidCharacterError("token contains whitespace")
		}
		if !slices.Contains(tokenList, token) {
			tokenList = append(tokenList, token)
		}
	}
	l.setTokens(tokenList)
	return nil
}

func (l DOMTokenList) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		tokens := l.getTokens()
		for _, token := range tokens {
			if !yield(token) {
				return
			}
		}
	}
}

func (l DOMTokenList) Contains(token string) bool {
	return slices.Contains(l.getTokens(), token)
}

func (l DOMTokenList) Length() int {
	return len(l.getTokens())
}

func (l DOMTokenList) Value() string {
	a, _ := l.element.GetAttribute(l.attribute)
	return a
}

func (l DOMTokenList) SetValue(val string) {
	l.element.SetAttribute(l.attribute, val)
}

func (l DOMTokenList) Item(index int) *string {
	tokens := l.getTokens()
	if index >= len(tokens) {
		return nil
	}
	return &tokens[index]
}

func (l DOMTokenList) Remove(token string) {
	tokens := l.getTokens()
	itemIndex := slices.Index(tokens, token)
	if itemIndex >= 0 {
		newList := slices.Delete(tokens, itemIndex, itemIndex+1)
		l.setTokens(newList)
	}
}

func (l DOMTokenList) Replace(oldToken string, newToken string) bool {
	if l.Contains(oldToken) {
		l.Remove(oldToken)
		l.Add(newToken)
		return true
	} else {
		return false
	}
}

func (l DOMTokenList) Toggle(token string) bool {
	if l.Contains(token) {
		l.Remove(token)
		return false
	} else {
		l.Add(token)
		return true
	}
}

func (l DOMTokenList) getTokens() []string {
	tokens := l.Value()
	if strings.TrimSpace(tokens) == "" {
		return []string{}
	}
	return strings.Split(tokens, " ")
}

func (l DOMTokenList) setTokens(tokens []string) {
	l.SetValue(strings.Join(tokens, " "))
}
