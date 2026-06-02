wpt-suite/MANIFEST.json:
	cd wpt-suite && ./wpt manifest

.PHONY: test ci
test:
	GOEXPERIMENT=jsonv2 go run . --log-type pretty \
				 dom/nodes/Document-getElementById \
				 dom/nodes/Element-childElement \
				 dom/nodes/Element-insert \
				 dom/nodes/CharacterData-appendChild

ci:
	./ci

# WIP target is used to run tests for currently worked on feature
.PHONY: wip
wip:
	GOEXPERIMENT=jsonv2 go run . --log-type pretty --log-level $${LOG_LEVEL:-warn} dom/nodes/DOMImplementation-createHTMLDocument.html
