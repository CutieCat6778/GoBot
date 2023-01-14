package class

import (
	"cutiecat6778/discordbot/commands"
	"time"
)

type (
	ForecastScrollStruct struct {
		UserID   string
		Position int
		Datas    map[string][]commands.ForecastMapStruct
		Dates    []int
	}

	ForecastScrollMap map[string]ForecastScrollStruct

	ForecastScroll struct {
		userscroll ForecastScrollMap
	}
)

func NewForecastScroll() *ForecastScroll {
	return &ForecastScroll{make(ForecastScrollMap)}
}

func (handler ForecastScroll) GetForecastScrolls() ForecastScrollMap {
	return handler.userscroll
}

func (handler ForecastScroll) Get(id string) (*ForecastScrollStruct, bool) {
	userscroll, found := handler.userscroll[id]

	return &userscroll, found
}

func (handler ForecastScroll) Write(id string, num int, dates []int) (*ForecastScrollStruct, bool) {
	userscroll, found := handler.userscroll[id]
	if !found {
		return &userscroll, found
	} else {
		userscroll.Position = num
		return &userscroll, found
	}
}

func (handler ForecastScroll) Register(id string, location map[string][]commands.ForecastMapStruct, dates []int) {
	userscrollstruct := ForecastScrollStruct{UserID: id, Position: time.Now().Day(), Datas: location}
	handler.userscroll[id] = userscrollstruct
	if len(id) > 1 {
		handler.userscroll[id[:1]] = userscrollstruct
	}
}

func (handler ForecastScrollStruct) GetPosition() int {
	return handler.Position
}

func (handler ForecastScrollStruct) GetLocation() map[string][]commands.ForecastMapStruct {
	return handler.Datas
}

func (handler ForecastScrollStruct) GetDates() []int {
	return handler.Dates
}
