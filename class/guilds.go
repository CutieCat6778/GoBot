package class

import "time"

type Guilds struct {
	ServerID  string `bson:"server_id,omitempty"`
	CreatedAt int64  `bson:"created_at,omitempty"`
}

func NewGuild(id string) Guilds {
	now := time.Now()
	guild := Guilds{id, now.Unix()}

	return guild
}
