package main

import (
	"encoding/json"
	"errors"
	"iter"
	"maps"
	"path"
	"slices"
	"sort"
	"strings"
)

type ManifestItems map[string]ManifestItem

type ItemKind int

const (
	KindUnknown ItemKind = iota
	KindItems
	KindItem
)

type WPTTest []any

type ManifestItem struct {
	Kind  ItemKind
	Items ManifestItems
	Item  WPTTest
}

func (i ManifestItem) All(p string) iter.Seq[TestCase] {
	return func(yield func(TestCase) bool) {
		switch i.Kind {
		case KindItem:
			if !yield(TestCase{p}) {
				return
			}
		case KindItems:
			keys := slices.Collect(maps.Keys(i.Items))
			sort.Strings(keys)
			for _, key := range keys {
				item := i.Items[key]
				for testCase := range item.All(path.Join(p, key)) {
					if !yield(testCase) {
						return
					}
				}
			}
		}
	}
}

func (i *ManifestItem) UnmarshalJSON(data []byte) error {
	err1 := json.Unmarshal(data, &i.Item)
	if err1 == nil {
		i.Kind = KindItem
		return nil
	}
	err2 := json.Unmarshal(data, &i.Items)
	if err2 == nil {
		i.Kind = KindItems
		return nil
	}
	return errors.Join(err1, err2)
}

type ManifestRootItems struct {
	Testharness map[string]ManifestItem `json:"testharness"`
}

type Manifest struct {
	Items   ManifestRootItems `json:"items"`
	UrlBase string            `json:"url_base"`
	Version int
}

type TestCase struct {
	Path string
}

func (m Manifest) All() iter.Seq[TestCase] {
	return func(yield func(TestCase) bool) {
		keys := slices.Collect(maps.Keys(m.Items.Testharness))
		sort.Strings(keys)
		for _, api := range keys {
			manifestItem := m.Items.Testharness[api]
			for testCase := range manifestItem.All(api) {
				switch strings.ToLower(path.Ext(testCase.Path)) {
				case ".htm", ".html":
					if !yield(testCase) {
						return
					}
				}
			}
		}
	}
}
