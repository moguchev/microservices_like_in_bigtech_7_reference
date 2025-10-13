package chat

import (
	"chat/internal/app/models"
	"chat/internal/app/models/types"
	chat_models "chat/internal/app/usecases/chat/models"
	pb "chat/pkg/api/chat/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func modelsChatToPb(c *models.Chat) *pb.Chat {
	if c == nil {
		return nil
	}

	participantIds := make([]string, len(c.Members))
	for i, member := range c.Members {
		participantIds[i] = member.String()
	}

	return &pb.Chat{
		ChatId:         c.ID.String(),
		ParticipantIds: participantIds,
	}
}

func modelsChatsToPb(chats []*models.Chat) []*pb.Chat {
	result := make([]*pb.Chat, len(chats))
	for i, c := range chats {
		result[i] = modelsChatToPb(c)
	}
	return result
}

func modelsMessageToPb(m *models.Message) *pb.Message {
	if m == nil {
		return nil
	}

	return &pb.Message{
		MessageId: m.ID.String(),
		ChatId:    m.ChatID.String(),
		SenderId:  m.SenderID.String(),
		Text:      m.Text,
		CreatedAt: timestamppb.New(m.CreatedAt),
	}
}

func modelsMessagesToPb(messages []*models.Message) []*pb.Message {
	result := make([]*pb.Message, len(messages))
	for i, m := range messages {
		result[i] = modelsMessageToPb(m)
	}
	return result
}

func userIDsToStrings(ids []types.UserID) []string {
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = id.String()
	}
	return result
}

func pbListMessagesRequestToModels(req *pb.ListMessagesRequest) *chat_models.ListMessagesRequest {
	return &chat_models.ListMessagesRequest{
		ChatID:          types.ChatID(req.GetChatId()),
		Limit:           req.GetLimit(),
		LastMessageTime: req.GetLastMessageTime().AsTime().UTC(),
	}
}
