package chat_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	"chat/internal/app/usecases/chat"
	"chat/internal/app/usecases/chat/mocks"
	chat_models "chat/internal/app/usecases/chat/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ListMessagesTestSuite struct {
	suite.Suite

	ChatRepository *mocks.ChatRepository
	UserIDProvider *mocks.UserIDProvider

	Usecase chat.Usecase
}

func (s *ListMessagesTestSuite) SetupTest() {
	s.ChatRepository = mocks.NewChatRepository(s.T())
	s.UserIDProvider = mocks.NewUserIDProvider(s.T())

	s.Usecase = chat.NewUsecase(chat.Deps{
		ChatRepository: s.ChatRepository,
		UserIDProvider: s.UserIDProvider,
	})
}

func (s *ListMessagesTestSuite) TearDownTest() {
	s.ChatRepository.AssertExpectations(s.T())
	s.UserIDProvider.AssertExpectations(s.T())
}

func TestListMessagesTestSuite(t *testing.T) {
	suite.Run(t, new(ListMessagesTestSuite))
}

func (s *ListMessagesTestSuite) Test_ListMessages_Positive() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	limit := uint32(10)
	lastTime := time.Now().UTC()

	chatFromRepo := &models.Chat{
		ID:      "chat1",
		Members: []types.UserID{"user1", "user2"},
	}

	expectedMessages := []*models.Message{
		{ID: "msg1", ChatID: "chat1", SenderID: "user2", Text: "hi"},
	}
	nextTime := lastTime.Add(-time.Minute)
	expectedRes := &chat_models.ListMessagesResult{
		Messages:            expectedMessages,
		NextLastMessageTime: nextTime,
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chatFromRepo, nil).
		Once()

	s.ChatRepository.EXPECT().
		GetMessages(
			ctx,
			chatID,
			mock.AnythingOfType("models.GetMessagesOption"),
			mock.AnythingOfType("models.GetMessagesOption"),
		).
		Return(&chat_models.GetMessagesResult{
			Messages:            expectedMessages,
			NextLastMessageTime: nextTime,
		}, nil).
		Once()

	gotRes, err := s.Usecase.ListMessages(ctx, &chat_models.ListMessagesRequest{
		ChatID:          chatID,
		Limit:           limit,
		LastMessageTime: lastTime,
	})

	s.NoError(err)
	s.Equal(expectedRes, gotRes)
}

func (s *ListMessagesTestSuite) Test_ListMessages_Unauthenticated() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	limit := uint32(10)
	lastTime := time.Now().UTC()

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("", errors.New("no user")).
		Once()

	gotRes, err := s.Usecase.ListMessages(ctx, &chat_models.ListMessagesRequest{
		ChatID:          chatID,
		Limit:           limit,
		LastMessageTime: lastTime,
	})

	s.Nil(gotRes)
	s.ErrorIs(err, models.ErrUnauthenticated)
}

func (s *ListMessagesTestSuite) Test_ListMessages_PermissionDenied() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	limit := uint32(10)
	lastTime := time.Now().UTC()

	chatFromRepo := &models.Chat{
		ID:      "chat1",
		Members: []types.UserID{"user2", "user3"},
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chatFromRepo, nil).
		Once()

	gotRes, err := s.Usecase.ListMessages(ctx, &chat_models.ListMessagesRequest{
		ChatID:          chatID,
		Limit:           limit,
		LastMessageTime: lastTime,
	})

	s.Nil(gotRes)
	s.ErrorIs(err, models.ErrPermissionDenied)
}
