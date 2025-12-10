.PHONY: test
test:
	GOEXPERIMENT=jsonv2 go run . dom/nodes/Element-childElement dom/nodes/Document-getElementById

# WIP target is used to run tests for currently worked on feature
.PHONY: wip
wip:
	GOEXPERIMENT=jsonv2 go run . --file dom/nodes/Element-insertAdjacentText.html
