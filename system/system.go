package system

type System struct {
	players    [maxPlayers]*Player
	actionMaps []*ActionMap
}

func NewSystem(players [maxPlayers]*Player, actionMaps []*ActionMap) *System {
	return &System{
		players:    players,
		actionMaps: actionMaps,
	}
}
