package social

import (
	"social/internal/app/models"
	"social/internal/app/models/types"
)

func fromModel(m *models.FriendRequest) *friendRequestRow {
	if m == nil {
		return nil
	}

	return &friendRequestRow{
		ID:        m.ID.String(),
		FromUser:  m.FromUser.String(),
		ToUser:    m.ToUser.String(),
		Status:    m.Status.String(),
		CreatedAt: m.CreatedAt,
	}
}

func toModel(r *friendRequestRow) *models.FriendRequest {
	if r == nil {
		return nil
	}

	return &models.FriendRequest{
		ID:        types.RequestID(r.ID),
		FromUser:  types.UserID(r.FromUser),
		ToUser:    types.UserID(r.ToUser),
		Status:    models.FriendRequestStatus(r.Status),
		CreatedAt: r.CreatedAt,
	}
}
