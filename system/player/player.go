package player

import (
	"github.com/google/uuid"

	"github.com/go-glx/input/system/controller"
)

type ID string

type Player struct {
	id  ID
	ctl *controller.Controller // can be nil
}

func NewPlayer() *Player {
	return &Player{
		id:  ID(uuid.NewString()),
		ctl: nil,
	}
}

func (p *Player) ID() ID {
	return p.id
}

func (p *Player) Controller() *controller.Controller {
	return p.ctl
}

func (p *Player) AttachController(ctl *controller.Controller) {
	p.ctl = ctl
}

func (p *Player) DetachController() {
	p.ctl = nil
}
