package exp

type Schema struct {
	layoutsVec1   []*Layout[Vec1]
	layoutsVec2   []*Layout[Vec2]
	layoutsAction []*Layout[Action]
	layoutsMouse  []*Layout[Mouse]

	lastState map[any]any
}

type SchemaInitializer = func(*Schema)

func NewSchema(opts ...SchemaInitializer) *Schema {
	sch := &Schema{
		layoutsVec1:   make([]*Layout[Vec1], 0),
		layoutsVec2:   make([]*Layout[Vec2], 0),
		layoutsAction: make([]*Layout[Action], 0),
		layoutsMouse:  make([]*Layout[Mouse], 0),

		lastState: make(map[any]any),
	}

	for _, init := range opts {
		init(sch)
	}

	return sch
}

func (s *Schema) Calculate() {
	s.lastState = make(map[any]any)

	for _, lay := range s.layoutsVec1 {
		s.lastState[lay.inputName] = lay.resolve()
	}

	for _, lay := range s.layoutsVec2 {
		s.lastState[lay.inputName] = lay.resolve()
	}

	for _, lay := range s.layoutsAction {
		s.lastState[lay.inputName] = lay.resolve()
	}

	for _, lay := range s.layoutsMouse {
		s.lastState[lay.inputName] = lay.resolve()
	}
}

func Get[T outputType](s *Schema, name any) T {
	if state, exist := s.lastState[name]; exist {
		return state.(T)
	}

	return *new(T)
}

func WithLayoutVec1(lay *Layout[Vec1]) SchemaInitializer {
	return func(schema *Schema) {
		schema.layoutsVec1 = append(schema.layoutsVec1, lay)
	}
}

func WithLayoutVec2(lay *Layout[Vec2]) SchemaInitializer {
	return func(schema *Schema) {
		schema.layoutsVec2 = append(schema.layoutsVec2, lay)
	}
}

func WithLayoutAction(lay *Layout[Action]) SchemaInitializer {
	return func(schema *Schema) {
		schema.layoutsAction = append(schema.layoutsAction, lay)
	}
}

func WithLayoutMouse(lay *Layout[Mouse]) SchemaInitializer {
	return func(schema *Schema) {
		schema.layoutsMouse = append(schema.layoutsMouse, lay)
	}
}
