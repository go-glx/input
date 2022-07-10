package action

type ID string

type Action struct {
	id     ID
	parent *Map
}

func NewAction(parent *Map, id ID) *Action {
	return &Action{
		id:     id,
		parent: parent,
	}
}

func (a *Action) ID() ID {
	return a.id
}

func (a *Action) Parent() *Map {
	return a.parent
}
