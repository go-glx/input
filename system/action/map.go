package action

type MapID string

type Map struct {
	id MapID
}

func NewMap(ID string) *Map {
	return &Map{id: MapID(ID)}
}

func (m *Map) ID() MapID {
	return m.id
}
