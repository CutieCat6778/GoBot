package class

type (
	SeaLevelScrollStruct struct {
		UserID   string
		Position int
		Location string
	}

	SeaLevelScrollMap map[string]SeaLevelScrollStruct

	SeaLevelScroll struct {
		userscroll SeaLevelScrollMap
	}
)

func NewSeaLevelScroll() *SeaLevelScroll {
	return &SeaLevelScroll{make(SeaLevelScrollMap)}
}

func (handler SeaLevelScroll) GetSeaLevelScrolls() SeaLevelScrollMap {
	return handler.userscroll
}

func (handler SeaLevelScroll) Get(id string) (*SeaLevelScrollStruct, bool) {
	userscroll, found := handler.userscroll[id]

	return &userscroll, found
}

func (handler SeaLevelScroll) Write(id string, num int) (*SeaLevelScrollStruct, bool) {
	userscroll, found := handler.userscroll[id]
	if !found {
		return &userscroll, found
	} else {
		userscroll.Position = num
		return &userscroll, found
	}
}

func (handler SeaLevelScroll) Register(id string, location string) {
	userscrollstruct := SeaLevelScrollStruct{UserID: id, Position: 0, Location: location}
	handler.userscroll[id] = userscrollstruct
	if len(id) > 1 {
		handler.userscroll[id[:1]] = userscrollstruct
	}
}

func (handler SeaLevelScrollStruct) GetPosition() int {
	return handler.Position
}

func (handler SeaLevelScrollStruct) GetLocation() string {
	return handler.Location
}
