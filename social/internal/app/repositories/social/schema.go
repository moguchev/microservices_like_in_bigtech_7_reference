package social

import "time"

const (
	tableFriendRequests = "public.friend_requests"
	tableFriends        = "public.friends"
)

const (
	columnID        = "id"
	columnFromUser  = "from_user_id"
	columnToUser    = "to_user_id"
	columnStatus    = "status"
	columnCreatedAt = "created_at"

	columnUserID       = "user_id"
	columnFriendUserID = "friend_user_id"
)

var tableFriendRequestsColumns = []string{
	columnID,
	columnFromUser,
	columnToUser,
	columnStatus,
	columnCreatedAt,
}

var tableFriendsColumns = []string{
	columnUserID,
	columnFriendUserID,
	columnCreatedAt,
}

type friendRow struct {
	UserID       string    `db:"user_id"`
	FriendUserID string    `db:"friend_user_id"`
	CreatedAt    time.Time `db:"created_at"`
}

type friendRequestRow struct {
	ID        string    `db:"id"`
	FromUser  string    `db:"from_user_id"`
	ToUser    string    `db:"to_user_id"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *friendRequestRow) mapFields() map[string]any {
	return map[string]any{
		columnID:        r.ID,
		columnFromUser:  r.FromUser,
		columnToUser:    r.ToUser,
		columnStatus:    r.Status,
		columnCreatedAt: r.CreatedAt,
	}
}

func (r *friendRequestRow) Values(columns ...string) []any {
	fields := r.mapFields()
	values := make([]any, 0, len(columns))
	for _, c := range columns {
		if v, ok := fields[c]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
