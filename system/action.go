package system

type actionID string
type actionMapID string

type Action struct {
	ID     actionID
	mapID  actionMapID
	system *System
}

func newAction(parent *System, id actionID, mapID actionMapID) *Action {
	return &Action{
		ID:     id,
		mapID:  mapID,
		system: parent,
	}
}
