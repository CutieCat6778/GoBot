package class

import "time"

type Members struct {
	MemberID      string `bson:"member_id,omitempty"`
	CreatedAt     int64  `bson:"created_at,omitempty"`
	Tokens        int64  `bson:"tokens,omitempty"`
	LastRefreshed int64  `bson:"last_refreshed,omitempty"`
}

func NewMember(id string) Members {
	now := time.Now().Unix()
	member := Members{id, now, 20, now}

	return member
}
