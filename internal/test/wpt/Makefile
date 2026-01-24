.PHONY: test
test:
	GOEXPERIMENT=jsonv2 go run . --log-type pretty \
				 dom/nodes/Document-getElementById \
				 dom/nodes/Element-childElement \
				 dom/nodes/Element-insert \
				 dom/nodes/CharacterData-appendChild

# WIP target is used to run tests for currently worked on feature
.PHONY: wip
wip:
	GOEXPERIMENT=jsonv2 go run . --file dom/events/Event-dispatch-click.html --log-type pretty --log-level debug 
	# | sed -r 's/\x1B\[(;?[0-9]{1,3})+[mGK]//g'

cdata:
	GOEXPERIMENT=jsonv2 go run . --log-type pretty --log-level debug -- dom/nodes/CharacterData
