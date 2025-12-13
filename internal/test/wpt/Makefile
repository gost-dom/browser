.PHONY: test
test:
	GOEXPERIMENT=jsonv2 go run . \
				 dom/events/Event-dispatch-redispatch.html \
				 dom/nodes/Document-getElementById \
				 dom/nodes/Element-childElement \
				 dom/nodes/Element-insert 

# WIP target is used to run tests for currently worked on feature
.PHONY: wip
wip:
	GOEXPERIMENT=jsonv2 go run . dom/events/Event
