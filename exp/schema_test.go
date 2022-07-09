package exp

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-glx/input/exp/bridge/keyboard/custom"
	input2 "github.com/go-glx/input/exp/input"
)

type testPlrMovementType struct{}

var testKeyboard *input2.VirtualKeyboard
var testKeyboardDriver *custom.KeyboardDriver

func testGetFakeKeyboardDriver() *custom.KeyboardDriver {
	if testKeyboardDriver == nil {
		testKeyboardDriver = custom.NewKeyboardDriver(true)
	}

	return testKeyboardDriver
}

func testGetFakeKeyboard() *input2.VirtualKeyboard {
	if testKeyboard == nil {
		testKeyboard = input2.NewVirtualKeyboard(
			testGetFakeKeyboardDriver(),
		)
	}

	return testKeyboard
}

func testCreateSchema() *Schema {
	inputKeyboard := testGetFakeKeyboard()

	inputLayoutKeyboardWASD := NewKeyboardVec2(input2.KeyW, input2.KeyS, input2.KeyA, input2.KeyD, inputKeyboard)
	inputLayoutKeyboardArrows := NewKeyboardVec2(input2.KeyUp, input2.KeyDown, input2.KeyLeft, input2.KeyRight, inputKeyboard)

	plrMovement := NewLayout[Vec2](
		testPlrMovementType{},
		inputLayoutKeyboardWASD,
		inputLayoutKeyboardArrows,
	)

	return NewSchema(
		WithLayoutVec2(plrMovement),
	)
}

func TestSchema_Calculate(t *testing.T) {
	fakeKeyboard := testGetFakeKeyboardDriver()
	schema := testCreateSchema()

	// zero input
	schema.Calculate()
	assert.Equal(t, Vec2{0.0, 0.0}, Get[Vec2](schema, testPlrMovementType{}))

	// change some inputs
	fakeKeyboard.SetState(input2.KeyW, custom.StateJustPressed)
	schema.Calculate()
	assert.Equal(t, Vec2{0.0, -1.0}, Get[Vec2](schema, testPlrMovementType{}))

	// change some inputs
	fakeKeyboard.SetState(input2.KeyW, custom.StateDown)
	schema.Calculate()
	assert.Equal(t, Vec2{0.0, -1.0}, Get[Vec2](schema, testPlrMovementType{}))

	// change some inputs
	fakeKeyboard.SetState(input2.KeyW, custom.StateReleased)
	fakeKeyboard.SetState(input2.KeyA, custom.StateJustPressed)
	schema.Calculate()
	assert.Equal(t, Vec2{-1.0, 0.0}, Get[Vec2](schema, testPlrMovementType{}))

	// change some inputs
	fakeKeyboard.SetState(input2.KeyW, custom.StateUp)
	fakeKeyboard.SetState(input2.KeyA, custom.StateUp)
	schema.Calculate()
	assert.Equal(t, Vec2{0.0, 0.0}, Get[Vec2](schema, testPlrMovementType{}))
}
