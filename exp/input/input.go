package input

type Kind string
type Name string

const (
	KindMouse          Kind = "mouse"
	KindKeyboard       Kind = "keyboard"
	KindGamepadGeneric Kind = "gamepad-generic"
	KindGamepadXbox    Kind = "gamepad-xbox"
	KindGamepadPS      Kind = "gamepad-ps"
	KindOther          Kind = "other"
)

type Input interface {
	Name() Name
	Kind() Kind
	IsEnabled() bool
	OnEnabled(func())
	OnDisabled(func())
}
