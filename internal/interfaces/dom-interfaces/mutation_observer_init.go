package dominterfaces

type ObserveOption = func(*MutationObserverInit)

type MutationObserverInit struct {
	Subtree               bool
	ChildList             bool
	Attributes            bool
	AttributeFilter       []string
	AttributeOldValue     bool
	CharacterData         bool
	CharacterDataOldValue bool
}
