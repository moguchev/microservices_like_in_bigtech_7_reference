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

type StreamMessagesTestSuite struct {
	suite.Suite

	ChatRepository *mocks.ChatRepository
	UserIDProvider *mocks.UserIDProvider

	Usecase chat.Usecase
}

func (s *StreamMessagesTestSuite) SetupTest() {
	s.ChatRepository = mocks.NewChatRepository(s.T())
	s.UserIDProvider = mocks.NewUserIDProvider(s.T())

	s.Usecase = chat.NewUsecase(chat.Deps{
		ChatRepository: s.ChatRepository,
		UserIDProvider: s.UserIDProvider,
	})
}

func (s *StreamMessagesTestSuite) TearDownTest() {
	s.ChatRepository.AssertExpectations(s.T())
	s.UserIDProvider.AssertExpectations(s.T())
}

func TestStreamMessagesTestSuite(t *testing.T) {
	suite.Run(t, new(StreamMessagesTestSuite))
}

func (s *StreamMessagesTestSuite) Test_StreamMessages_Positive() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chatID := types.ChatID("chat1")
	since := time.Now().Add(-time.Minute)

	chat := &models.Chat{
		ID:      chatID,
		Members: []types.UserID{"user1", "user2"},
	}

	historyMessages := []*models.Message{
		{ID: "m1", ChatID: chatID, SenderID: "user2", Text: "hi"},
	}

	streamCh := make(chan *models.Message, 1)
	streamCh <- &models.Message{ID: "m2", ChatID: chatID, SenderID: "user2", Text: "hello"}
	close(streamCh)

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chat, nil).
		Once()

	s.ChatRepository.EXPECT().
		GetMessages(
			ctx,
			chatID,
			mock.AnythingOfType("models.GetMessagesOption"),
			mock.AnythingOfType("models.GetMessagesOption"),
		).
		Return(&chat_models.GetMessagesResult{
			Messages:            historyMessages,
			NextLastMessageTime: since,
		}, nil).
		Once()

	s.ChatRepository.EXPECT().
		StreamMessages(ctx, chatID).
		Return(streamCh, nil).
		Once()

	msgCh, err := s.Usecase.StreamMessages(ctx, chatID, since)
	s.NoError(err)

	var got []*models.Message
	for m := range msgCh {
		got = append(got, m)
	}

	s.Len(got, 2)
	s.Equal(types.MessageID("m1"), got[0].ID)
	s.Equal(types.MessageID("m2"), got[1].ID)
}

func (s *StreamMessagesTestSuite) Test_StreamMessages_WithHistory() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chatID := types.ChatID("chat1")
	since := time.Now().Add(-time.Hour)

	chat := &models.Chat{
		ID:      chatID,
		Members: []types.UserID{"user1", "user2"},
	}

	historyMessages := []*models.Message{
		{ID: "m1", ChatID: chatID, SenderID: "user2", Text: "history1"},
		{ID: "m2", ChatID: chatID, SenderID: "user2", Text: "history2"},
	}

	streamCh := make(chan *models.Message)
	close(streamCh)

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chat, nil).
		Once()

	s.ChatRepository.EXPECT().
		GetMessages(
			ctx,
			chatID,
			mock.AnythingOfType("models.GetMessagesOption"),
			mock.AnythingOfType("models.GetMessagesOption"),
		).
		Return(&chat_models.GetMessagesResult{
			Messages:            historyMessages,
			NextLastMessageTime: since,
		}, nil).
		Once()

	s.ChatRepository.EXPECT().
		StreamMessages(ctx, chatID).
		Return(streamCh, nil).
		Once()

	msgCh, err := s.Usecase.StreamMessages(ctx, chatID, since)
	s.NoError(err)

	var got []*models.Message
	for m := range msgCh {
		got = append(got, m)
	}

	s.Len(got, 2)
	s.Equal(types.MessageID("m1"), got[0].ID)
	s.Equal(types.MessageID("m2"), got[1].ID)
}

func (s *StreamMessagesTestSuite) Test_StreamMessages_Unauthenticated() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("", errors.New("no user")).
		Once()

	msgCh, err := s.Usecase.StreamMessages(ctx, chatID, time.Now())

	s.Nil(msgCh)
	s.ErrorIs(err, models.ErrUnauthenticated)
}

func (s *StreamMessagesTestSuite) Test_StreamMessages_PermissionDenied() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")

	chat := &models.Chat{
		ID:      chatID,
		Members: []types.UserID{"user2", "user3"},
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chat, nil).
		Once()

	msgCh, err := s.Usecase.StreamMessages(ctx, chatID, time.Now())

	s.Nil(msgCh)
	s.ErrorIs(err, models.ErrPermissionDenied)
}
