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

func validateDomToken(token string) error {
	if token == "" {
		return newSyntaxError("token is empty")
	}
	if strings.Contains(token, " ") {
		return newInvalidCharacterError("token contains whitespace")
	}
	return nil
}

func (l DOMTokenList) Add(tokens ...string) error {
	tokenList := l.getTokens()
	for _, token := range tokens {
		if err := validateDomToken(token); err != nil {
			return err
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

func (l DOMTokenList) Item(index int) (string, bool) {
	tokens := l.getTokens()
	if index >= len(tokens) || index < 0 {
		return "", false
	}
	return tokens[index], true
}

// Remove implements DOMTokenList.Remove
//
// see also: https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/remove
func (l DOMTokenList) Remove(token ...string) error {
	tokens := l.getTokens()
	for _, t := range token {
		if err := validateDomToken(t); err != nil {
			return err
		}
		itemIndex := slices.Index(tokens, t)
		if itemIndex >= 0 {
			tokens = slices.Delete(tokens, itemIndex, itemIndex+1)
		}

	}
	l.setTokens(tokens)
	return nil
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

func (l DOMTokenList) Toggle(token string) (bool, error) {
	if l.Contains(token) {
		l.Remove(token)
		return false, nil
	} else {
		err := l.Add(token)
		return true, err
	}
}

func (l DOMTokenList) getTokens() []string {
	tokens := strings.TrimSpace(l.Value())
	if tokens == "" {
		return []string{}
	}
	return strings.Fields(tokens)
}

func (l DOMTokenList) setTokens(tokens []string) {
	l.SetValue(strings.Join(tokens, " "))
}
