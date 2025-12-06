package main

import "strings"

type pathFilter struct {
	// included specifies the subpaths of the "testharness" tests that will be
	// included.
	included []string
	// excluded filters individual subpaths that would have been included from
	// the includeList.
	excluded []string
}

func (f pathFilter) isMatch(tc TestCase) bool {
	path := tc.Path
	for _, include := range f.included {
		if strings.HasPrefix(path, include) {
			for _, exclude := range f.excluded {
				if strings.HasPrefix(path, exclude) {
					return false
				}
			}

			return true
		}
	}
	return false

}
