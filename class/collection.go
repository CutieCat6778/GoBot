package class

import "time"

type (
	RatelimitStruct struct {
		id   string
		last int64
	}

	RatelimitMap map[string]RatelimitStruct

	Ratelimit struct {
		ratelimit RatelimitMap
	}
)

func NewRatelimit() *Ratelimit {
	return &Ratelimit{make(RatelimitMap)}
}

func (handler Ratelimit) GetRatelimits() RatelimitMap {
	return handler.ratelimit
}

func (handler Ratelimit) Get(id string) (*RatelimitStruct, bool) {
	ratelimit, found := handler.ratelimit[id]

	return &ratelimit, found
}

func (handler Ratelimit) Write(id string) (*RatelimitStruct, bool) {
	now := time.Now()
	ratelimit, found := handler.ratelimit[id]
	if !found {
		return &ratelimit, found
	} else {
		ratelimit.last = now.Unix()
		return &ratelimit, found
	}
}

func (handler Ratelimit) Register(id string) {
	now := time.Now()
	ratelimitstruct := RatelimitStruct{id: id, last: now.Unix()}
	handler.ratelimit[id] = ratelimitstruct
	if len(id) > 1 {
		handler.ratelimit[id[:1]] = ratelimitstruct
	}
}

func (ratelimit RatelimitStruct) GetTime() int {
	return int(ratelimit.last)
}
