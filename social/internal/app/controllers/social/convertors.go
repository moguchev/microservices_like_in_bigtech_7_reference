package social

import (
	"social/internal/app/models"
	"social/internal/app/models/types"
	social_models "social/internal/app/usecases/social/models"
	pb "social/pkg/api/social/v1"
)

func modelsFriendRequestToPb(fr *models.FriendRequest) *pb.FriendRequest {
	if fr == nil {
		return nil
	}
	return &pb.FriendRequest{
		RequestId: fr.ID.String(),
		Status:    friendRequestStatusToPb(fr.Status),
	}
}

func modelsFriendRequestsToPb(requests []*models.FriendRequest) []*pb.FriendRequest {
	result := make([]*pb.FriendRequest, len(requests))
	for i, fr := range requests {
		result[i] = modelsFriendRequestToPb(fr)
	}
	return result
}

func friendRequestStatusToPb(status models.FriendRequestStatus) pb.FriendRequestStatus {
	switch status {
	case models.FriendRequestStatusPending:
		return pb.FriendRequestStatus_FRIEND_REQUEST_STATUS_PENDING
	case models.FriendRequestStatusAccepted:
		return pb.FriendRequestStatus_FRIEND_REQUEST_STATUS_ACCEPTED
	case models.FriendRequestStatusDeclined:
		return pb.FriendRequestStatus_FRIEND_REQUEST_STATUS_DECLINED
	default:
		return pb.FriendRequestStatus_FRIEND_REQUEST_STATUS_UNSPECIFIED
	}
}

func userIDsToStrings(ids []types.UserID) []string {
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = id.String()
	}
	return result
}

func pbListFriendsRequestToModels(req *pb.ListFriendsRequest) *social_models.ListFriendsRequest {
	return &social_models.ListFriendsRequest{
		UserID:                types.UserID(req.GetUserId()),
		Limit:                 req.GetLimit(),
		LastAcceptRequestTime: req.GetLastAcceptRequestTime().AsTime().UTC(),
	}
}
